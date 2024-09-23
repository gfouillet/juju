# **D**e**L**ve every**W**here

This module allows to wrap any module easily into a delve server, enabling easy debugging, even for remote process.

It ships the delve dependencies with the binary, so it shouldn't used in production, and you should use custom tags to
build debug binaries.

To make it works:

1. compile your code with `-gcflags "all=-N -l"`
2. Use `dlw.Wrap` function to encapsulate your main function.

# Tips

Don't bring delve in production:

Add next to your `main.go` a `init_debug.go` file which will be compiled only if the tags debug is
passed to the compiler:

```sh
go build -gcflags "all=-N -l" -tags debug path/to/my/package
```

This file will contain an init function which will do wrap the main function into Delve,
only if the tag is defined. This way, delve dependencies won't be shipped in production.

# Example

## [init_debug.go](examples/init_debug.go)

```go
//go:build debug
// +build debug

package main

import (
	"github.com/juju/juju/internal/dlw"
	"github.com/juju/juju/internal/dlw/config"
)

func init() {
	runMain = dlw.Wrap(config.Default())(runMain)
}
```

## [main.go](examples/main.go)

```go
package main

import (
	"os"
)

var runMain = mainArgs

func main() {
	os.Exit(runMain(os.Args))
}

func mainArgs(args []string) int { /* ... */ }
```