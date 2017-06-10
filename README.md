# Color

Package `color` provides a facility to colorize text.

## Install

To install, use `go get`:

```bash
$ go get -u github.com/elpinal/color
```

-u flag stands for "update".

## Examples

```go
package main

import (
	"fmt"
	"os"

	"github.com/elpinal/color"
)

func main() {
	w := color.New(os.Stdout, color.Blue)
	fmt.Fprint(w, "Hello")
}
```

## Contribution

1. Fork ([https://github.com/elpinal/color/fork](https://github.com/elpinal/color/fork))
1. Create a feature branch
1. Commit your changes
1. Rebase your local changes against the master branch
1. Run test suite with the `go test ./...` command and confirm that it passes
1. Run `gofmt -s`
1. Create a new Pull Request

## Author

[elpinal](https://github.com/elpinal)
