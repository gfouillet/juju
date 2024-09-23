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
	// TODO(gfouillet): Find a way to not have to call /home/<USER>/go/bin/juju to make it works.
	commands.Main = dlw.Wrap(config.Default(),
		config.WithPort(10120),
		config.WaitDebugger(),
	)(commands.Main)
}
