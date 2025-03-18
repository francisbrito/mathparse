package lexer

import (
	"mathparse/token"
	"unicode"
)

type Lexer struct {
	input string
	pos   int
	ch    rune
}

func New(input string) *Lexer {
	return &Lexer{input: input, pos: -1}
}

func (l *Lexer) NextToken() token.Token {
	var literal string
	var tokenType token.Type
	l.readChar()
	l.skipWhitespace()
	switch l.ch {
	case 0:
		literal = ""
		tokenType = token.EndOfFile
	default:
		// tokenize integer
		if unicode.IsDigit(l.ch) {
			literal = l.tokenizeInteger()
			tokenType = token.Integer
		} else {
			// any other token is considered illegal
			literal = l.input[l.pos:]
			tokenType = token.Illegal
		}
	}
	return newToken(tokenType, literal)
}

func (l *Lexer) readChar() {
	l.pos++
	if l.pos == len(l.input) {
		l.ch = 0
	} else {
		l.ch = rune(l.input[l.pos])
	}
}

func (l *Lexer) skipWhitespace() {
	for unicode.IsSpace(l.ch) {
		l.readChar()
	}
}

func (l *Lexer) tokenizeInteger() string {
	start := l.pos
	for unicode.IsDigit(l.ch) {
		l.readChar()
	}
	return l.input[start:l.pos]
}

func newToken(tokenType token.Type, literal string) token.Token {
	return token.Token{Type: tokenType, Literal: literal}
}
