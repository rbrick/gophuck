// Echoes back the arguments to stdout
package main

import (
	"github.com/rbrick/gophuck"
	"os"
	"strings"
)

func main() {
	i := gophuck.NewInterpreter(",[.,]", strings.Join(os.Args[1:], " "), os.Stdout)
	i.Interpret()
}
