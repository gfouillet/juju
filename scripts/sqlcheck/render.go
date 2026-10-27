// Copyright 2026 Canonical Ltd.
// Licensed under the AGPLv3, see LICENCE file for details.

package main

import (
	"sort"
	"unicode/utf8"
)

func renderPackage(pkg *parsedPackage, stmts []Statement) map[string]*FileRender {
	byFile := map[string][]Statement{}
	for _, st := range stmts {
		if st.Eval.empty() || st.File == "" {
			continue
		}
		byFile[st.File] = append(byFile[st.File], st)
	}
	out := map[string]*FileRender{}
	for path, list := range byFile {
		pf := pkg.files[path]
		if pf == nil {
			continue
		}
		sort.Slice(list, func(i, j int) bool {
			if list[i].Start != list[j].Start {
				return list[i].Start < list[j].Start
			}
			return list[i].ID < list[j].ID
		})
		out[path] = renderFile(pf, list)
	}
	return out
}

func renderFile(pf *parsedFile, stmts []Statement) *FileRender {
	r := &FileRender{GoPath: pf.path}
	goPos := 0
	for _, st := range stmts {
		if st.Start >= goPos {
			r.emitGap(pf.src[goPos:st.Start])
			goPos = st.End
			if goPos < st.Start {
				goPos = st.Start
			}
		} else {
			r.emitVirtual([]byte("\n"))
		}
		r.emitStatement(st)
	}
	if goPos < len(pf.src) {
		r.emitGap(pf.src[goPos:])
	}
	if len(r.Synthetic) == 0 || r.Synthetic[len(r.Synthetic)-1] != '\n' {
		r.emitVirtual([]byte("\n"))
	}
	return r
}

func (r *FileRender) emitStatement(st Statement) {
	ev := st.Eval
	if ev.empty() {
		return
	}
	end := len(ev.Content)
	for end > 0 {
		switch ev.Content[end-1] {
		case '\n', '\r':
			end--
		default:
			goto split
		}
	}
split:
	if end > 0 {
		r.emitEval(&evaluated{Content: ev.Content[:end], Map: ev.Map[:end]}, st.ID)
	}
	if !endsWithSemi(ev.Content[:end]) {
		r.emitVirtual([]byte(";"))
	}
	if end < len(ev.Content) {
		r.emitEval(&evaluated{Content: ev.Content[end:], Map: ev.Map[end:]}, st.ID)
	}
}

func endsWithSemi(s string) bool {
	for i := len(s) - 1; i >= 0; i-- {
		switch s[i] {
		case ' ', '\t', '\n', '\r':
			continue
		case ';':
			return true
		default:
			return false
		}
	}
	return false
}

func (r *FileRender) emitGap(src []byte) {
	n := 0
	for _, b := range src {
		if b == '\n' {
			n++
		}
	}
	if n == 0 {
		return
	}
	r.emitVirtual([]byte(nNewlines(n)))
}

func nNewlines(n int) string {
	b := make([]byte, n)
	for i := range b {
		b[i] = '\n'
	}
	return string(b)
}

func (r *FileRender) emitVirtual(b []byte) {
	if len(b) == 0 {
		return
	}
	start := len(r.Synthetic)
	r.Synthetic = append(r.Synthetic, b...)
	r.Segments = append(r.Segments, Segment{
		RenderedStart: start,
		RenderedEnd:   len(r.Synthetic),
		Kind:          KindVirtual,
	})
}

func (r *FileRender) emitEval(ev *evaluated, stmtID int) {
	if ev.empty() {
		return
	}
	runStart := 0
	for i := 1; i <= len(ev.Content); i++ {
		if i < len(ev.Content) && sameRun(ev.Map[i-1], ev.Map[i]) {
			continue
		}
		r.emitRun(ev, runStart, i, stmtID)
		runStart = i
	}
}

func sameRun(a, b origByte) bool {
	if a.Kind != b.Kind || a.File != b.File {
		return false
	}
	switch a.Kind {
	case KindReplacement, KindPadding, KindEscape, KindVirtual:
		return a.Offset == b.Offset
	default:
		return b.Offset == a.Offset+1
	}
}

func (r *FileRender) emitRun(ev *evaluated, start, end, stmtID int) {
	if start >= end {
		return
	}
	off := len(r.Synthetic)
	r.Synthetic = append(r.Synthetic, ev.Content[start:end]...)
	ob := ev.Map[start]
	last := ev.Map[end-1]
	origEnd := last.Offset + 1
	if ob.Kind == KindReplacement || ob.Kind == KindPadding || ob.Kind == KindEscape {
		origEnd = ob.Offset + 1
	}
	r.Segments = append(r.Segments, Segment{
		RenderedStart: off,
		RenderedEnd:   len(r.Synthetic),
		OriginalFile:  ob.File,
		OriginalStart: ob.Offset,
		OriginalEnd:   origEnd,
		Kind:          ob.Kind,
		StmtID:        stmtID,
	})
}

func lineColToOffset(src []byte, line, col int) int {
	if line < 1 {
		line = 1
	}
	if col < 1 {
		col = 1
	}
	cur := 1
	i := 0
	for cur < line && i < len(src) {
		if src[i] == '\n' {
			cur++
		}
		i++
	}
	runes := 0
	for runes < col-1 && i < len(src) && src[i] != '\n' {
		_, size := utf8.DecodeRune(src[i:])
		if size < 1 {
			size = 1
		}
		i += size
		runes++
	}
	return i
}
