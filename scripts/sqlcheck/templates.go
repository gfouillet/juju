// Copyright 2026 Canonical Ltd.
// Licensed under the AGPLv3, see LICENCE file for details.

package main

import (
	"io/fs"
	"os"
	"path/filepath"
	"strings"
)

func checkTemplates(root string) []TemplateNote {
	var notes []TemplateNote
	genDir := filepath.Join(root, "generate")
	info, err := os.Stat(genDir)
	if err != nil || !info.IsDir() {
		return nil
	}
	_ = filepath.WalkDir(genDir, func(path string, d fs.DirEntry, err error) error {
		if err != nil || d.IsDir() {
			return err
		}
		ext := filepath.Ext(path)
		if ext != ".tmpl" && ext != ".gotmpl" {
			return nil
		}
		data, err := os.ReadFile(path)
		if err != nil {
			return nil
		}
		if !templateLooksLikeSQL(string(data)) {
			return nil
		}
		rel, err := filepath.Rel(root, path)
		if err != nil {
			rel = path
		}
		rel = filepath.ToSlash(rel)
		if strings.HasPrefix(rel, "generate/export/") {
			if fileExists(filepath.Join(root, "domain/export/state/model/export.go")) {
				return nil
			}
			notes = append(notes, TemplateNote{
				Path:    rel,
				Message: "SQL-bearing export template has no committed generated Go output",
			})
			return nil
		}
		notes = append(notes, TemplateNote{
			Path:    rel,
			Message: "SQL-bearing template is not linted directly; generated Go output was not found",
		})
		return nil
	})
	return notes
}

func templateLooksLikeSQL(s string) bool {
	upper := strings.ToUpper(s)
	for _, kw := range []string{"SELECT ", "INSERT ", "UPDATE ", "DELETE ", "CREATE ", "ALTER "} {
		if strings.Contains(upper, kw) {
			return true
		}
	}
	return strings.Contains(s, "sqlair.Prepare") || strings.Contains(s, "MakePatch")
}

func fileExists(path string) bool {
	_, err := os.Stat(path)
	return err == nil
}
