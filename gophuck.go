// TODO: Remove redundant code .-.
// TODO: Where `pos++` occurs: Make Command a struct and add something like `proceed bool`, then after interpret increase
// TODO(cont'd): CONT: if needed (if c.proceed { pos++ })
package gophuck

import (
	"fmt"
	"io"
	"log"
	"strings"
)

type Command rune

const (
	IncrementPointer Command = '>'
	DecrementPointer Command = '<'
	BeginLoop        Command = '['
	EndLoop          Command = ']'
	OuputByte        Command = '.'
	AcceptByte       Command = ','
	IncrementValue   Command = '+'
	DecrementValue   Command = '-'
)

type Interpreter interface {
	// Interpret interprets Brainfuck source code
	Interpret()
}

// implementation of Interpreter
type interpreter struct {
	input   strings.Reader
	output  io.Writer
	pointer int
	cells   []rune

	pos   int
	insts []Command
}

func (i interpreter) Interpret() {
	for i.pos < len(i.insts) {
		c := i.insts[i.pos]
		switch c {
		case IncrementPointer:
			i.pointer++
			i.pos++
		case DecrementPointer:
			{
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
				i.pos++
				break
			}
			if err != nil {
				log.Panicln(err)
				break
			}
			i.cells[i.pointer] = r
			i.pos++
		case BeginLoop:
			//If the byte at the data pointer is zero (cells[pointer] == 0) jump forward to command after matching ']'
			matchingEnd := i.findMatchingEnd()
			if matchingEnd == -1 {
				log.Panic("could not find matching end brace")
			}

			if i.cells[i.pointer] == 0 {
				i.pos = matchingEnd + 1
			} else {
				i.pos++
			}
		case EndLoop:
			// If the byte at the data pointer is nonzero (cells[pointer] > 0) jump back to the command after matching '['
			matchingBegin := i.findMatchingBegin()
			if matchingBegin == -1 {
				log.Panic("could not find matching begin brace")
			}

			if i.cells[i.pointer] > 0 {
				i.pos = matchingBegin + 1
			} else {
				i.pos++
			}
		default:
			i.pos++
			continue
		}
	}
}

/*
   how findMatching* works.
   It will look through and find all the matching open/close commands.
   If it finds a match (open to close), it will decrement the match count (default: 1)
   unless it is equal to the default value (1). If it reaches the default value, we found our match.

   e.g.
   findMatchingEnd
   [[]]

   Searching for end
   count = 1

   [ - skip first
   [ - finds this, increment count
   count = 2
   ] - Sees closing brace, decrement
   count = 1
   ] - Sees closing brace
   return
 */

// Finds the position of the matching end loop
func (i *interpreter) findMatchingEnd() int {
	res := -1
	openBraces := 1
	initial := i.pos
	for x := i.pos; x < len(i.insts); x++ {
		if x == initial {
			continue
		}
		if i.insts[x] == BeginLoop {
			openBraces++
		} else if i.insts[x] == EndLoop {
			if openBraces == 1 {
				res = x
				break
			} else {
				openBraces--
			}
		}
	}
	return res
}

func (i *interpreter) findMatchingBegin() int {
	res := -1
	closedBraces := 1
	initial := i.pos
	// We move back
	for x := i.pos; x > 0; x-- {
		if x == initial {
			continue
		}
		if i.insts[x] == EndLoop {
			// At closed brace
			closedBraces++
		} else if i.insts[x] == BeginLoop {
			// At begin
			if closedBraces == 1 {
				res = x
				break
			} else {
				closedBraces--
			}
		}
	}
	return res
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

func NewInterpreter(source string, input string, output io.Writer) Interpreter {
	i := interpreter{}
	i.input = *strings.NewReader(input)
	i.output = output
	i.pointer = 0
	i.cells = make([]rune, 1024)

	// Instructions
	i.pos = 0
	i.insts = parseSource(source)
	return i
}
