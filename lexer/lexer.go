package lexer

import "mathparse/token"

type Lexer struct {
	input string
}

func New(input string) *Lexer {
	return &Lexer{input: input}
}

func (l *Lexer) NextToken() token.Token {
	return token.Token{Type: token.EndOfFile}
}
