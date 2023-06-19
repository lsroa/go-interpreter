package main

import (
	"fmt"
	"lsroa/go_interpreter/repl"
	"os"
)

func main() {
	fmt.Println("Monkey language:")
	repl.Start(os.Stdin, os.Stdout)
}
