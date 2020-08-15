package main

import (
	"bufio"
	"fmt"
	"io"
	"os"

	"github.com/adzeitor/stopka"
)

func scan(scanner *bufio.Scanner) bool {
	fmt.Print("> ")
	return scanner.Scan()
}

func repl(input io.Reader, output io.Writer) {
	machine := stopka.New()
	buf := bufio.NewScanner(input)
	for scan(buf) {
		line := buf.Text()
		machine.Eval(line)
		fmt.Fprintf(output, "%+v\n", machine.Stack())
		if machine.IsHalted() {
			fmt.Fprintf(output, "exception: %v\n\n", machine.Err)
		}
		fmt.Fprintln(output)
	}
}

func main() {
	repl(os.Stdin, os.Stdout)
}
