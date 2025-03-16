package lexer

import (
	"unicode"
)

type TokenType string

const (
	Integer   TokenType = "INTEGER"
	Float               = "FLOAT"
	Unknown             = "UNKNOWN"
	EndOfFile           = "EOF"
	MinusSign           = "MINUS_SIGN"
)

type Token struct {
	Type    TokenType
	Literal string
}

type Lexer struct {
	s   string
	pos int
}

func isDigit(ch uint8) bool {
	return unicode.IsDigit(rune(ch))
}

func (l *Lexer) tokenizeInteger() {
	sLen := len(l.s)
	for l.pos < sLen {
		ch := l.s[l.pos]
		if isDigit(ch) {
			l.pos++
		} else {
			break
		}
	}
}

func (l *Lexer) tokenizeFloatingPoint() {

}

func (l *Lexer) Tokenize() []Token {
	var parsed []Token
	sLen := len(l.s)
	isNegative := false
	for l.pos < sLen {
		ch := l.s[l.pos]
		switch {
		// parse integers and floating point numbers.
		case ch == '-':
			// if number is already negative, this is an error condition.
			if isNegative {
				parsed = []Token{{Unknown, l.s}}
				l.pos = sLen
				break
			}
			if l.pos+1 < sLen && isDigit(l.s[l.pos+1]) {
				isNegative = true
				token := Token{MinusSign, "-"}
				parsed = append(parsed, token)
			} else {
				parsed = []Token{{Unknown, l.s}}
				l.pos = sLen
				break
			}
		case isDigit(ch):
			start := l.pos
			l.tokenizeInteger()
			token := Token{Integer, l.s[start:l.pos]}
			parsed = append(parsed, token)
			continue
		default:
			token := Token{Unknown, string(ch)}
			parsed = append(parsed, token)
		}
		l.pos++
	}
	return append(parsed, Token{EndOfFile, ""})
}

func New(s string) *Lexer {
	return &Lexer{s: s}
}
