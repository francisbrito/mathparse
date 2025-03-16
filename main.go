package main

import (
	"fmt"
	"mathparse/lexer"
)

func main() {
	l := lexer.New("2 + 2")
	tokens := l.Tokenize()
	fmt.Printf("tokens=%v\n", tokens)
}
