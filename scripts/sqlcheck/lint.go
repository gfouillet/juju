// Copyright 2026 Canonical Ltd.
// Licensed under the AGPLv3, see LICENCE file for details.

package main

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"sort"
	"strings"
)

type fluffFile struct {
	Filepath   string           `json:"filepath"`
	Violations []fluffViolation `json:"violations"`
}

type fluffViolation struct {
	LineNo       int    `json:"line_no"`
	LinePos      int    `json:"line_pos"`
	StartLineNo  int    `json:"start_line_no"`
	StartLinePos int    `json:"start_line_pos"`
	Code         string `json:"code"`
	Description  string `json:"description"`
	Name         string `json:"name"`
}

func (v fluffViolation) line() int {
	if v.StartLineNo != 0 {
		return v.StartLineNo
	}
	return v.LineNo
}

func (v fluffViolation) col() int {
	if v.StartLinePos != 0 {
		return v.StartLinePos
	}
	return v.LinePos
}

func (v fluffViolation) rule() string {
	if v.Code != "" {
		return v.Code
	}
	return v.Name
}

func writeExtracted(dir string, renders map[string]*FileRender) (map[string]*FileRender, error) {
	if err := os.MkdirAll(dir, 0o755); err != nil {
		return nil, err
	}
	bySynth := map[string]*FileRender{}
	for path, rend := range renders {
		if len(rend.Synthetic) == 0 {
			continue
		}
		synth := filepath.Join(dir, filepath.FromSlash(path)+".sql")
		if err := os.MkdirAll(filepath.Dir(synth), 0o755); err != nil {
			return nil, err
		}
		if err := os.WriteFile(synth, rend.Synthetic, 0o644); err != nil {
			return nil, err
		}
		bySynth[synth] = rend
	}
	return bySynth, nil
}

func runSQLFluff(ctx context.Context, bin, config, dir string, rules map[string]bool) ([]fluffFile, string, error) {
	args := []string{"lint", "--config", config, "--format", "json"}
	if len(rules) > 0 {
		args = append(args, "--rules", joinRules(rules))
	}
	args = append(args, dir)
	cmd := exec.CommandContext(ctx, bin, args...)
	var stdout, stderr bytes.Buffer
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr
	err := cmd.Run()
	if err != nil && !isFluffLintExit(err) {
		return nil, stderr.String(), fmt.Errorf("sqlfluff lint failed: %w\n%s", err, stderr.String())
	}
	data := bytes.TrimSpace(stdout.Bytes())
	if len(data) == 0 {
		return nil, stderr.String(), nil
	}
	var files []fluffFile
	if err := json.Unmarshal(data, &files); err != nil {
		return nil, stderr.String(), fmt.Errorf("parsing sqlfluff JSON: %w\n%s", err, truncate(string(data), 2000))
	}
	return files, stderr.String(), nil
}

func joinRules(rules map[string]bool) string {
	names := make([]string, 0, len(rules))
	for r := range rules {
		names = append(names, r)
	}
	sort.Strings(names)
	return strings.Join(names, ",")
}

func isFluffLintExit(err error) bool {
	ex, ok := err.(*exec.ExitError)
	if !ok {
		return false
	}
	code := ex.ExitCode()
	return code == 1 || code == 0
}

func lookSQLFluff() (string, error) {
	bin, err := exec.LookPath("sqlfluff")
	if err != nil {
		return "", fmt.Errorf("sqlfluff not found in PATH; install sqlfluff==3.0.7")
	}
	return bin, nil
}

func matchSynth(bySynth map[string]*FileRender, fluffPath string) *FileRender {
	if r, ok := bySynth[fluffPath]; ok {
		return r
	}
	cleaned := filepath.Clean(fluffPath)
	if r, ok := bySynth[cleaned]; ok {
		return r
	}
	for synth, r := range bySynth {
		if synth == fluffPath || filepath.Base(synth) == filepath.Base(fluffPath) {
			return r
		}
		if strings.HasSuffix(filepath.ToSlash(fluffPath), filepath.ToSlash(synth)) {
			return r
		}
		if strings.HasSuffix(filepath.ToSlash(synth), filepath.ToSlash(fluffPath)) {
			return r
		}
	}
	return nil
}

func truncate(s string, n int) string {
	if len(s) <= n {
		return s
	}
	return s[:n] + "..."
}
