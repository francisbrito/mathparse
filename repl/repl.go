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
		parseErrors := p.Errors()
		if expression == nil {
			if len(parseErrors) > 0 {
				for _, parserError := range parseErrors {
					fmt.Println(parserError)
				}
			}
		} else {
			fmt.Println(expression.String())
		}
		fmt.Println()
	}
}
