// Copyright 2012, 2013 Canonical Ltd.
// Licensed under the AGPLv3, see LICENCE file for details.

//go:build debug

package main

import (
	"github.com/juju/juju/internal/dlw"
	"github.com/juju/juju/internal/dlw/config"
)

func init() {
	Main = dlw.Wrap(config.Default(),
		config.WithPort(10122),
		config.NoWait(),
	)(Main)
}
