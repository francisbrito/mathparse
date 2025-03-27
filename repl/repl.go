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
		expression := p.Parse()
		for _, expr := range expression.Terms {
			fmt.Print(expr.String())
		}
		if parserErrors := p.Errors(); len(parserErrors) > 0 {
			fmt.Println()
			for _, parserError := range parserErrors {
				fmt.Println(parserError)
			}
		}
		fmt.Println()
		fmt.Println()
	}
}
