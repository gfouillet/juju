// Copyright 2024 Canonical Ltd.
// Licensed under the AGPLv3, see LICENCE file for details.

//go:build debug

package main

import (
	"github.com/juju/juju/internal/dlw"
	"github.com/juju/juju/internal/dlw/config"

	"github.com/juju/juju/cmd/juju/commands"
)

func init() {
	commands.Main = dlw.Wrap(config.Default(),
		config.WithPort(10121),
		config.WaitDebugger(),
	)(commands.Main)
}
