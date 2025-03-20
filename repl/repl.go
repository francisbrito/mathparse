package repl

import (
	"bufio"
	"fmt"
	"io"
	"mathparse/lexer"
	"mathparse/token"
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
		for tok := l.NextToken(); tok.Type != token.EndOfFile; tok = l.NextToken() {
			fmt.Printf("%+v\n", tok)
		}
		fmt.Println()
	}
}
