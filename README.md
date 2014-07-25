#completecommand

Programmatic autocompletion suggestions for Go binaries.

## Description

`completecommand` is a Go library that allows your binaries to programmatically return a list of
accepted flags that can be used to provide shell autocompletion suggestions for command line
invocations.

## Usage

In your program, simply import `completecommand`, and call its `Complete` functions as the first
line of your `main`.

```go
package main

import (
        "flag"

        "github.com/tiziano88/completecommand"
)

var flagA = flag.Bool("a", false, "A")
...

func main() {
        completecommand.Complete()
        ...
}
```

Then create a file defining an autocomletion function based on the output of running your executable
with the `__complete__` flag, and `source` it in your `.zshrc`.

```zsh
func _exec {
  reply=($(exec -__complete__))
}

compctl -K _exec exec
```
