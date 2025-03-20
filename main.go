package main

import (
	"mathparse/repl"
	"os"
)

func main() {
	repl.Start(os.Stdin, os.Stdout)
}
