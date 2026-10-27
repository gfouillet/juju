// Copyright 2026 Canonical Ltd.
// Licensed under the AGPLv3, see LICENCE file for details.

package main

import (
	"context"
	"go/token"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/juju/tc"
)

type integrationSuite struct{}

func TestIntegrationSuite(t *testing.T) {
	tc.Run(t, &integrationSuite{})
}

func (s *integrationSuite) TestSQLFluffEndToEnd(c *tc.C) {
	bin, err := lookSQLFluff()
	if err != nil {
		if os.Getenv("CI") != "" {
			c.Fatalf("sqlfluff is required in CI: %v", err)
		}
		c.Skip("sqlfluff not installed")
	}

	// sqlfluff 3.0.7 with .github/.sqlfluff does not emit CP01 for
	// lowercase keywords. LT01 (layout.spacing) is a known hit.
	src := trimFixture(`
package p

func f() {
	st.Prepare(` + "`" + `
SELECT id  FROM t
WHERE x = 1
` + "`" + `)
}
`)
	fset := token.NewFileSet()
	pkg, err := parsePackage(fset, []goFile{{Path: "foo.go", Src: []byte(src)}})
	c.Assert(err, tc.ErrorIsNil)
	next := 1
	stmts, skips := extractPackage(pkg, &next)
	c.Assert(skips, tc.HasLen, 0)
	c.Assert(stmts, tc.HasLen, 1)
	rend := renderPackage(pkg, stmts)["foo.go"]
	c.Assert(rend, tc.NotNil)

	root, err := findRepoRoot(".")
	c.Assert(err, tc.ErrorIsNil)
	cfg := filepath.Join(root, ".github", ".sqlfluff")

	dir := c.MkDir()
	bySynth, err := writeExtracted(dir, map[string]*FileRender{"foo.go": rend})
	c.Assert(err, tc.ErrorIsNil)

	absSynth := map[string]*FileRender{}
	for p, r := range bySynth {
		abs, _ := filepath.Abs(p)
		absSynth[abs] = r
		absSynth[p] = r
	}

	fluff, stderr, err := runSQLFluff(context.Background(), bin, cfg, dir)
	c.Assert(err, tc.ErrorIsNil, tc.Commentf("stderr: %s", stderr))
	diags := mapDiagnostics(pkg.fset, stmts, absSynth, fluff)
	c.Assert(len(diags) > 0, tc.IsTrue)
	found := false
	for _, d := range diags {
		if d.Rule == "LT01" && d.Line == 5 {
			found = true
			c.Check(d.File, tc.Equals, "foo.go")
			c.Check(d.Column > 0, tc.IsTrue)
			c.Check(strings.Contains(strings.ToLower(d.Message), "space"), tc.IsTrue)
		}
	}
	c.Check(found, tc.IsTrue, tc.Commentf("diagnostics: %#v", diags))
}
