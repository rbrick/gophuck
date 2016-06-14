// Specify a source file to run. It will use passed in arguments as input
package main

import (
	"flag"
	"io/ioutil"
	"log"
	"github.com/rbrick/gophuck"
	"os"
	"strings"
)


func main()  {
        var file string

	flag.StringVar(&file, "f", "source.b", "source file")

	flag.Parse()

	d, err := ioutil.ReadFile(file)
	if err != nil {
		log.Fatal(err.Error())
	}

	interpreter := gophuck.NewInterpreter(string(d), strings.Join(flag.Args(), " "), os.Stdout)
	interpreter.Interpret()
}