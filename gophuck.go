package main

import (
	"strings"
	"io"
	"log"
	"fmt"
)

type Command rune

const (
	MovePointerRight Command = '>'
	MovePointerLeft Command = '<'
	BeginLoop Command = '['
	EndLoop Command = ']'
	PutChar Command = '.'
	GetChar Command = ','
	Add Command = '+'
	Sub Command = '-'
)

func Parse(src, input string) {
	cells := make([]rune, 1024)
	pos := 0

	reader := strings.NewReader(src)
	inputReader := strings.NewReader(input)
	for {
		r, _, err := reader.ReadRune()
		if err == io.EOF {
			break
		}

		if err != nil {
			log.Panicln(err)
			break
		}

		switch Command(r) {
		case MovePointerRight:
			pos++
		case MovePointerLeft: {
			if pos > 0 {
				pos--
			}
		}
		case Add:
			cells[pos] += 1
		case Sub:
			cells[pos] -= 1
		case PutChar:
			fmt.Print(string(cells[pos]))
		case GetChar:
			r, _, err := inputReader.ReadRune()
			if err != nil {
				log.Panicln(err)
				break
			}
			cells[pos] = r
		default:
			continue
		}
	}
}

func main() {
	// Output is 1234
	Parse(",+.>,++.>,+++.>,++++.", "0000")
}