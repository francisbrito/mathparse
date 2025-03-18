package main

import (
	"bufio"
	"fmt"
	"mathparse/lexer"
	"mathparse/token"
	"os"
)

func main() {
	fmt.Print("> ")
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		l := lexer.New(scanner.Text())
		fmt.Printf(">>> ")
		for {
			tok := l.NextToken()
			if tok.Type == token.EndOfFile {
				break
			} else {
				fmt.Printf("%v ", tok)
				if tok.Type == token.Illegal {
					break
				}
			}
		}
		fmt.Printf("\n\n> ")
	}
}
