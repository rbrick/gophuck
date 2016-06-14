// Prints "Hello, World!"
package main

import (
	"github.com/rbrick/gophuck"
	"os"
)

func main() {
	// Hello World!
	interpreter := gophuck.NewInterpreter(">++++++++[-<+++++++++>]<.>>+>-[+]++>++>+++[>[->+++<<+++>]<<]>-----.>->+++..+++.>-.<<+[>[+>+]>>]<--------------.>>.+++.------.--------.>+.>+.", "", os.Stdout)
	interpreter.Interpret()
}
