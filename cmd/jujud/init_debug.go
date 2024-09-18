//go:build debug

package main

import (
	"github.com/fogfactory/dlw"
	"github.com/fogfactory/dlw/config"
)

func init() {
	Main = dlw.Wrap(config.Default(),
		config.WithPort(10122),
		config.WaitDebugger(),
	)(Main)
}
