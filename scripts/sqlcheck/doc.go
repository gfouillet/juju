// Copyright 2026 Canonical Ltd.
// Licensed under the AGPLv3, see LICENCE file for details.

// Command sqlcheck extracts SQL embedded in Juju's Go source, renders it
// into SQLite-compatible SQL for the repository's pinned sqlfluff, and
// reports any linting diagnostics against the original Go file, line,
// and column.
//
// SQL that cannot be statically reconstructed, such as imperative string
// composition, is reported in a coverage summary instead of being
// silently skipped. Use `make lint-sql` from the repository root.
package main
