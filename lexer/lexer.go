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
	l := &Lexer{input: input, pos: -1}
	l.readChar()
	return l
}

func (l *Lexer) NextToken() token.Token {
	var literal string
	var tokenType token.Type
	l.skipWhitespace()
	switch l.ch {
	case 0:
		literal = ""
		tokenType = token.EndOfFile
	case '+':
		literal = "+"
		tokenType = token.PlusSign
		l.readChar()
	case '-':
		literal = "-"
		tokenType = token.MinusSign
		l.readChar()
	case '*':
		literal = "*"
		tokenType = token.Asterisk
		l.readChar()
	case '/':
		literal = "/"
		tokenType = token.Slash
		l.readChar()
	case '%':
		literal = "%"
		tokenType = token.PercentSign
		l.readChar()
	case '^':
		literal = "^"
		tokenType = token.Caret
		l.readChar()
	case '(':
		literal = "("
		tokenType = token.OpeningParentheses
		l.readChar()
	case ')':
		literal = ")"
		tokenType = token.ClosingParentheses
		l.readChar()
	default:
		if unicode.IsDigit(l.ch) {
			literal, tokenType = l.tokenizeNumber()
		} else {
			// any other token is considered illegal
			start := l.pos
			for !unicode.IsSpace(l.ch) && l.ch != 0 {
				l.readChar()
			}
			literal = l.input[start:l.pos]
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

func (l *Lexer) tokenizeNumber() (string, token.Type) {
	var isFloat bool
	tokenType := token.Integer
	start := l.pos
	for unicode.IsDigit(l.ch) || l.ch == '.' {
		if l.ch == '.' {
			if isFloat {
				tokenType = token.Illegal
			} else {
				isFloat = true
				tokenType = token.Float
			}
		}
		l.readChar()
	}
	return l.input[start:l.pos], token.Type(tokenType)
}

func newToken(tokenType token.Type, literal string) token.Token {
	return token.Token{Type: tokenType, Literal: literal}
}
