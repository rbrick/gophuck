package main

import (
	"github.com/rbrick/gophuck"
	"strings"
	"os"
)

func main() {
	input := strings.NewReader("0000")
	interpreter := gophuck.NewInterpreter(",+.>,++.>,+++.>,++++.", input, os.Stdout)
	interpreter.Interpret()
}
