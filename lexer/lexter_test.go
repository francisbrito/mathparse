package lexer

import (
	"github.com/stretchr/testify/assert"
	"mathparse/token"
	"testing"
)

func TestLexer_NextToken(t *testing.T) {
	l := New("")
	assert.Equal(t, token.Token{Type: token.EndOfFile}, l.NextToken())
}
