package lexer

import (
	"github.com/stretchr/testify/assert"
	"mathparse/token"
	"testing"
)

const digits = "1234567890"

func TestLexer_NextToken(t *testing.T) {
	var l *Lexer

	// can process empty input
	l = New("")
	assert.Equal(t, token.Token{Type: token.EndOfFile}, l.NextToken())

	// can tokenize illegal tokens
	l = New("ILLEGAL")
	assert.Equal(t, token.Token{Type: token.Illegal, Literal: "ILLEGAL"}, l.NextToken())

	// can skip whitespace
	l = New(" ")
	assert.Equal(t, token.Token{Type: token.EndOfFile}, l.NextToken())
	l = New("		")
	assert.Equal(t, token.Token{Type: token.EndOfFile}, l.NextToken())

	// can tokenize every digit
	for _, digit := range digits {
		l = New(string(digit))
		if !assert.Equal(t, token.Token{Type: token.Integer, Literal: string(digit)}, l.NextToken()) {
			t.Fatal()
		}
	}

	// can tokenize integers
	l = New("1234567890")
	assert.Equal(t, token.Token{Type: token.Integer, Literal: "1234567890"}, l.NextToken())
}
