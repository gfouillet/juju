// Copyright 2026 Canonical Ltd.
// Licensed under the AGPLv3, see LICENCE file for details.

// Command sqlcheck extracts SQL from Go source, lints it with sqlfluff,
// and reports diagnostics against the original Go positions.
package main

import (
	"context"
	"flag"
	"fmt"
	"go/token"
	"os"
	"os/signal"
	"path/filepath"
	"syscall"
)

func main() {
	os.Exit(run(os.Args[1:]))
}

func run(args []string) int {
	fs := flag.NewFlagSet("sqlcheck", flag.ContinueOnError)
	fs.SetOutput(os.Stderr)
	reportOnly := fs.Bool("report-only", false, "print violations but exit 0 unless the tool itself fails")
	github := fs.Bool("github", false, "emit GitHub Actions error annotations")
	writeDir := fs.String("write-extracted", "", "write rendered SQL into this directory instead of a temp dir")
	config := fs.String("config", "", "path to sqlfluff config (default <repo>/.github/.sqlfluff)")
	fs.Usage = func() {
		fmt.Fprintf(fs.Output(), "usage: sqlcheck [flags] [files-or-dirs...]\n")
		fs.PrintDefaults()
	}
	if err := fs.Parse(args); err != nil {
		if err == flag.ErrHelp {
			return 0
		}
		return 2
	}

	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer stop()

	cwd, err := os.Getwd()
	if err != nil {
		fmt.Fprintf(os.Stderr, "sqlcheck: %v\n", err)
		return 2
	}
	root, err := findRepoRoot(cwd)
	if err != nil {
		fmt.Fprintf(os.Stderr, "sqlcheck: %v\n", err)
		return 2
	}

	if os.Getenv("GITHUB_ACTIONS") == "true" {
		*github = true
	}

	cfg := *config
	if cfg == "" {
		cfg = filepath.Join(root, ".github", ".sqlfluff")
	}
	if !filepath.IsAbs(cfg) {
		cfg = filepath.Join(cwd, cfg)
	}
	if _, err := os.Stat(cfg); err != nil {
		fmt.Fprintf(os.Stderr, "sqlcheck: sqlfluff config not found: %s\n", cfg)
		return 2
	}

	bin, err := lookSQLFluff()
	if err != nil {
		fmt.Fprintf(os.Stderr, "sqlcheck: %v\n", err)
		return 2
	}

	rels, err := discoverGoFiles(root, fs.Args())
	if err != nil {
		fmt.Fprintf(os.Stderr, "sqlcheck: %v\n", err)
		return 2
	}
	files, err := readGoFiles(root, rels)
	if err != nil {
		fmt.Fprintf(os.Stderr, "sqlcheck: %v\n", err)
		return 2
	}

	fset := token.NewFileSet()
	var stmts []Statement
	var skips []Skip
	renders := map[string]*FileRender{}
	nextID := 1
	for _, group := range groupByPackage(files) {
		pkg, err := parsePackage(fset, group)
		if err != nil {
			fmt.Fprintf(os.Stderr, "sqlcheck: %v\n", err)
			return 2
		}
		st, sk := extractPackage(pkg, &nextID)
		stmts = append(stmts, st...)
		skips = append(skips, sk...)
		for path, rend := range renderPackage(pkg, st) {
			renders[path] = rend
		}
	}

	notes := checkTemplates(root)

	extractDir := *writeDir
	cleanup := false
	if extractDir == "" {
		tmp, err := os.MkdirTemp("", "sqlcheck-")
		if err != nil {
			fmt.Fprintf(os.Stderr, "sqlcheck: %v\n", err)
			return 2
		}
		extractDir = tmp
		cleanup = true
	}
	if cleanup {
		defer os.RemoveAll(extractDir)
	}

	bySynth, err := writeExtracted(extractDir, renders)
	if err != nil {
		fmt.Fprintf(os.Stderr, "sqlcheck: %v\n", err)
		return 2
	}

	absSynth := map[string]*FileRender{}
	for p, r := range bySynth {
		abs, err := filepath.Abs(p)
		if err != nil {
			abs = p
		}
		absSynth[abs] = r
		absSynth[p] = r
	}

	var fluff []fluffFile
	if len(bySynth) > 0 {
		var err error
		fluff, _, err = runSQLFluff(ctx, bin, cfg, extractDir)
		if err != nil {
			fmt.Fprintf(os.Stderr, "sqlcheck: %v\n", err)
			return 2
		}
	}

	diags := mapDiagnostics(fset, stmts, absSynth, fluff)
	for _, d := range diags {
		fmt.Println(formatDiagnostic(d, *github))
	}

	nFiles := 0
	seen := map[string]bool{}
	for _, st := range stmts {
		if !seen[st.File] {
			seen[st.File] = true
			nFiles++
		}
	}
	fmt.Fprint(os.Stderr, writeSummary(len(stmts), nFiles, skips, notes, len(diags)))

	if len(diags) > 0 && !*reportOnly {
		return 1
	}
	return 0
}
