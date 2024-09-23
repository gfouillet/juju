// Package dlw allows to run any binaries with latest version of delve.
//
// To make it works:
//  1. compile your code with
//     -gcflags "all=-N -l"
//  2. Use [Wrap] function to encapsulate your main function.
//
// # Tips
//
// Don't bring delve in production:
//
// Add next to your main.go a init_debug.go file which will be compiled only if the tags debug is
// passed to the compiler:
//
//	go build -gcflags "all=-N -l" -tags debug path/to/my/package
//
// This file will contain an init function which will wrap the main function into Delve,
// only if the tag is defined. This way, delve dependencies won't be shipped in production.
//
// # Example
//
// debug_init.go
//
//	//go:build debug
//
//	package main
//
//	 import (
//	 "github.com/juju/juju/internal/dlw"
//	 "github.com/juju/juju/internal/dlw/config"
//	 )
//
//	func init() {
//	   runMain = dlw.Wrap(config.Default())(runMain)
//	}
//
// main.go
//
//	package main
//
//	import (
//	   "os"
//	)
//
//	var runMain = mainArgs
//
//	func main() {
//	   os.Exit(runMain(os.Args))
//	}
//
//	func mainArgs(args []string) int { /* ... */ }
package dlw
