// Copyright 2024 Canonical Ltd.
// Licensed under the AGPLv3, see LICENCE file for details.

//go:build debug

package main

import (
	"github.com/fogfactory/dlw"
	"github.com/fogfactory/dlw/config"

	"github.com/juju/juju/cmd/juju/commands"
)

func init() {
	commands.Main = dlw.Wrap(config.Default(),
		config.WithPort(10122),
		config.WaitDebugger(),
	)(commands.Main)
}
