package repl

import (
	"bufio"
	"fmt"
	"io"
	"lsroa/go_interpreter/lexer"
	"lsroa/go_interpreter/token"
)

func Start(in io.Reader, out io.Writer) {
	scanner := bufio.NewScanner(in)
	for {
		fmt.Fprintf(out, ">> ")
		scanned := scanner.Scan()

		if !scanned {
			return
		}

		line := scanner.Text()
		if line == ".exit" {
			fmt.Fprintln(out, "Bye")
			return
		}
		l := lexer.New(line)
		for tok := l.NextToken(); tok.Kind != token.EOF; tok = l.NextToken() {
			fmt.Fprintf(out, "%+v\n", tok)
		}
	}
}
