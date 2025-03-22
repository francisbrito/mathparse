package repl

import (
	"bufio"
	"fmt"
	"io"
	"mathparse/lexer"
	"mathparse/parser"
)

const Prompt = ">>>"

func Start(in io.Reader, out io.Writer) {
	scanner := bufio.NewScanner(in)
	for {
		fmt.Printf("%s ", Prompt)
		scanned := scanner.Scan()
		if !scanned {
			return
		}
		l := lexer.New(scanner.Text())
		p := parser.New(l)
		expressions := p.Parse()
		for _, expr := range expressions {
			fmt.Print(expr.String())
		}
		fmt.Println()
	}
}
