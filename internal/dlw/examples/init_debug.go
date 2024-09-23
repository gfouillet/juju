//go:build debug
// +build debug

package main

import (
	"github.com/juju/juju/internal/dlw"
	"github.com/juju/juju/internal/dlw/config"
)

func init() {
	runMain = dlw.Wrap(config.Default(),
		config.WithPort(1122),
		config.WaitDebugger())(runMain)
}
