// TODO: Remove redundant code .-.
// TODO: Where `pos++` occurs: Make Command a struct and add something like `proceed bool`, then after interpret increase
// TODO(cont'd): CONT: if needed (if c.proceed { pos++ })
package gophuck

import (
	"strings"
	"io"
	"log"
	"fmt"
	"bufio"
)

type Command rune

const (
	IncrementPointer Command = '>'
	DecrementPointer Command = '<'
	BeginLoop Command = '['
	EndLoop Command = ']'
	OuputByte Command = '.'
	AcceptByte Command = ','
	IncrementValue Command = '+'
	DecrementValue Command = '-'
)

type Interpreter interface {
	// Interpret interprets Brainfuck source code
	Interpret()
}

// implementation of Interpreter
type interpreter struct {
	input   bufio.Reader
	output  io.Writer
	pointer int
	cells   []rune


	pos     int
	insts   []Command
}

func (i interpreter) Interpret() {
	for i.pos < len(i.insts) {
		c := i.insts[i.pos]
		switch c {
		case IncrementPointer:
			i.pointer++
			i.pos++
		case DecrementPointer: {
			if i.pointer > 0 {
				i.pointer--
			}
			i.pos++
		}
		case IncrementValue:
			i.cells[i.pointer]++
			i.pos++
		case DecrementValue:
			i.cells[i.pointer]--
			i.pos++
		case OuputByte:
			fmt.Print(string(i.cells[i.pointer]))
			i.pos++
		case AcceptByte:
			r, _, err := i.input.ReadRune()
			if err == io.EOF {
				i.cells[i.pointer] = 0
				break
			}
			if err != nil {
				log.Panicln(err)
				break
			}
			i.cells[i.pointer] = r
			i.pos++
		// case BeginLoop:
		// If the byte at the data pointer is zero (cells[pointer] == 0) jump forward to command after matching ']'
		// case EndLoop:
		// If the byte at the data pointer is nonzero (cells[pointer] > 0) jump back to the command after matching '['
		default:
			i.pos++
			continue
		}
	}
}

func parseSource(s string) []Command {
	source := strings.NewReader(s)
	insts := make([]Command, source.Len())
	for x := 0; x < len(insts); x++ {
		r, _, err := source.ReadRune()
		if err == io.EOF {
			break
		}

		if err != nil {
			log.Panicln(err)
			break
		}
		insts[x] = Command(r)
	}
	return insts
}

func NewInterpreter(source string, input io.Reader, output io.Writer) Interpreter {
	i := interpreter{}
	i.input = *bufio.NewReader(input)
	i.output = output
	i.pointer = 0
	i.cells = make([]rune, 1024)

	// Instructions
	i.pos = 0
	i.insts = parseSource(source)
	return i
}