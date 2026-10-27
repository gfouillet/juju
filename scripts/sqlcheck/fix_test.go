// Copyright 2026 Canonical Ltd.
// Licensed under the AGPLv3, see LICENCE file for details.

package main

import (
	"go/token"
	"strings"
	"testing"

	"github.com/juju/tc"
)

type fixSuite struct{}

func TestFixSuite(t *testing.T) {
	tc.Run(t, &fixSuite{})
}

// fixFixture extracts, renders and plans fixes for the given violations
// against a single-file fixture.
func fixFixture(c *tc.C, src string, violations []fluffViolation) ([]sourceEdit, int) {
	c.Helper()
	src = trimFixture(src)
	fset := token.NewFileSet()
	pkg, err := parsePackage(fset, []goFile{{Path: "foo.go", Src: []byte(src)}})
	c.Assert(err, tc.ErrorIsNil)
	next := 1
	stmts, skips := extractPackage(pkg, &next)
	c.Assert(skips, tc.HasLen, 0)
	c.Assert(stmts, tc.Not(tc.HasLen), 0)
	rend := renderPackage(pkg, stmts)["foo.go"]
	c.Assert(rend, tc.NotNil)
	edits, skipped := planFixes(stmts, map[string]*FileRender{
		"foo.go.sql": rend,
	}, []fluffFile{{
		Filepath:   "foo.go.sql",
		Violations: violations,
	}}, pkg.files)
	return edits, skipped
}

// violationAt builds a violation anchored at the first byte of needle in
// the rendered file.
func violationAt(c *tc.C, rend *FileRender, needle, code, desc string) fluffViolation {
	c.Helper()
	idx := strings.Index(string(rend.Synthetic), needle)
	c.Assert(idx >= 0, tc.IsTrue, tc.Commentf("needle %q not found in %q", needle, rend.Synthetic))
	line, col := 1, 1
	for i := 0; i < idx; i++ {
		if rend.Synthetic[i] == '\n' {
			line++
			col = 1
		} else {
			col++
		}
	}
	return fluffViolation{LineNo: line, LinePos: col, Code: code, Description: desc}
}

func renderFixture(c *tc.C, src string) (*FileRender, []Statement) {
	c.Helper()
	src = trimFixture(src)
	fset := token.NewFileSet()
	pkg, err := parsePackage(fset, []goFile{{Path: "foo.go", Src: []byte(src)}})
	c.Assert(err, tc.ErrorIsNil)
	next := 1
	stmts, skips := extractPackage(pkg, &next)
	c.Assert(skips, tc.HasLen, 0)
	rend := renderPackage(pkg, stmts)["foo.go"]
	c.Assert(rend, tc.NotNil)
	return rend, stmts
}

func applyTo(c *tc.C, src string, edits []sourceEdit) string {
	c.Helper()
	for _, e := range edits {
		c.Assert(e.file, tc.Equals, "foo.go")
		c.Assert(e.start >= 0 && e.end <= len(src) && e.start <= e.end, tc.IsTrue)
		src = src[:e.start] + e.repl + src[e.end:]
	}
	return src
}

func (s *fixSuite) TestCollapseSpace(c *tc.C) {
	src := `package p

func f() {
	st.Prepare("SELECT id FROM   t WHERE x = 1")
}
`
	rend, _ := renderFixture(c, src)
	v := violationAt(c, rend, "   t", "LT01",
		"Expected only single space before naked identifier. Found '   '.")
	edits, skipped := fixFixture(c, src, []fluffViolation{v})
	c.Assert(skipped, tc.Equals, 0)
	c.Assert(edits, tc.HasLen, 1)
	out := applyTo(c, src, edits)
	c.Check(strings.Contains(out, "FROM t WHERE"), tc.IsTrue)
}

func (s *fixSuite) TestRemoveTrailingWhitespace(c *tc.C) {
	src := "package p\n\nfunc f() {\n\tst.Prepare(`SELECT id\nFROM t \nWHERE x = 1`)\n}\n"
	rend, _ := renderFixture(c, src)
	v := violationAt(c, rend, " \nWHERE", "LT01", "Unnecessary trailing whitespace.")
	edits, skipped := fixFixture(c, src, []fluffViolation{v})
	c.Assert(skipped, tc.Equals, 0)
	out := applyTo(c, src, edits)
	c.Check(strings.Contains(out, "FROM t\nWHERE"), tc.IsTrue)
}

func (s *fixSuite) TestInsertSpaceBeforeBracket(c *tc.C) {
	src := `package p

func f() {
	st.Prepare("INSERT INTO t (a) VALUES (1) ON CONFLICT(a) DO NOTHING")
}
`
	rend, _ := renderFixture(c, src)
	v := violationAt(c, rend, "(a) DO NOTHING", "LT01",
		"Expected single whitespace between 'CONFLICT' keyword and start bracket '('.")
	edits, skipped := fixFixture(c, src, []fluffViolation{v})
	c.Assert(skipped, tc.Equals, 0)
	out := applyTo(c, src, edits)
	c.Check(strings.Contains(out, "ON CONFLICT (a)"), tc.IsTrue)
}

func (s *fixSuite) TestInsertSpaceAroundOperator(c *tc.C) {
	src := `package p

func f() {
	st.Prepare("UPDATE t SET name='x' WHERE id = 1")
}
`
	rend, _ := renderFixture(c, src)
	v := violationAt(c, rend, "='x'", "LT01",
		"Expected single whitespace between naked identifier and raw comparison operator '='.")
	edits, skipped := fixFixture(c, src, []fluffViolation{v})
	c.Assert(skipped, tc.Equals, 0)
	out := applyTo(c, src, edits)
	c.Check(strings.Contains(out, "name ="), tc.IsTrue)
}

func (s *fixSuite) TestIndentLine(c *tc.C) {
	src := "package p\n\nfunc f() {\n\tst.Prepare(`SELECT a FROM t\nAND b = 1`)\n}\n"
	rend, _ := renderFixture(c, src)
	v := violationAt(c, rend, "AND b", "LT02", "Expected indent of 4 spaces.")
	edits, skipped := fixFixture(c, src, []fluffViolation{v})
	c.Assert(skipped, tc.Equals, 0)
	out := applyTo(c, src, edits)
	c.Check(strings.Contains(out, "FROM t\n    AND b = 1"), tc.IsTrue)
}

func (s *fixSuite) TestDedentLine(c *tc.C) {
	src := "package p\n\nfunc f() {\n\tst.Prepare(`SELECT a\n    FROM t\nWHERE b = 1`)\n}\n"
	rend, _ := renderFixture(c, src)
	v := violationAt(c, rend, "    FROM t", "LT02", "Line should not be indented.")
	edits, skipped := fixFixture(c, src, []fluffViolation{v})
	c.Assert(skipped, tc.Equals, 0)
	out := applyTo(c, src, edits)
	c.Check(strings.Contains(out, "SELECT a\nFROM t\nWHERE"), tc.IsTrue)
}

func (s *fixSuite) TestLineBreakInRawString(c *tc.C) {
	// A newline insertion into a raw string literal stays a real
	// newline in the Go source.
	src := "package p\n\nfunc f() {\n\tst.Prepare(`SELECT a, b FROM t WHERE x = 1`)\n}\n"
	rend, _ := renderFixture(c, src)
	v := violationAt(c, rend, "a, b", "LT02",
		"Expected line break and indent of 4 spaces before 'a'.")
	edits, _ := fixFixture(c, src, []fluffViolation{v})
	out := applyTo(c, src, edits)
	c.Check(strings.Contains(out, "SELECT\n    a, b FROM"), tc.IsTrue)
}

func (s *fixSuite) TestBlankLineAfterCTE(c *tc.C) {
	src := "package p\n\nfunc f() {\n\tst.Prepare(`WITH x AS (\n    SELECT a\n)\nSELECT a FROM x WHERE a = 1`)\n}\n"
	rend, _ := renderFixture(c, src)
	v := violationAt(c, rend, "SELECT a FROM x", "LT08",
		"Blank line expected but not found after CTE closing bracket.")
	edits, skipped := fixFixture(c, src, []fluffViolation{v})
	c.Assert(skipped, tc.Equals, 0)
	out := applyTo(c, src, edits)
	c.Check(strings.Contains(out, ")\n\nSELECT a FROM x"), tc.IsTrue)
}

func (s *fixSuite) TestInsertionInInterpretedStringEscapesNewline(c *tc.C) {
	// A newline insertion into an interpreted string must insert the
	// two-byte \n escape, not a raw newline that would break the Go
	// source.
	src := "package p\n\nfunc f() {\n\tst.Prepare(\"SELECT a, b FROM t WHERE x = 1\")\n}\n"
	rend, _ := renderFixture(c, src)
	v := violationAt(c, rend, "a, b", "LT02",
		"Expected line break and indent of 4 spaces before 'a'.")
	edits, _ := fixFixture(c, src, []fluffViolation{v})
	out := applyTo(c, src, edits)
	c.Check(strings.Contains(out, `SELECT\n    a, b FROM`), tc.IsTrue)
	c.Check(strings.Contains(out, "\n"), tc.IsTrue)
}

func (s *fixSuite) TestSkipsVirtualAnchoredViolation(c *tc.C) {
	// A violation anchored in synthetic content (the blanked Go code
	// before the first statement) must not produce an edit.
	src := "package p\n\nfunc f() {\n\tst.Prepare(`SELECT a FROM t WHERE a = 1`)\n}\n"
	edits, skipped := fixFixture(c, src, []fluffViolation{{
		LineNo:      1,
		LinePos:     1,
		Code:        "LT01",
		Description: "Unnecessary trailing whitespace.",
	}})
	c.Check(len(edits), tc.Equals, 0)
	c.Check(skipped, tc.Equals, 1)
}

func (s *fixSuite) TestSkipsUnfixableMessage(c *tc.C) {
	src := "package p\n\nfunc f() {\n\tst.Prepare(`SELECT a FROM t WHERE a = 1`)\n}\n"
	edits, skipped := fixFixture(c, src, []fluffViolation{{
		LineNo:      4,
		LinePos:     14,
		Code:        "PRS",
		Description: "Line 4, Position 14: Found unparsable section",
	}})
	c.Check(len(edits), tc.Equals, 0)
	c.Check(skipped, tc.Equals, 1)
}
