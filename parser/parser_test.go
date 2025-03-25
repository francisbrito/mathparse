package parser

import (
	"github.com/stretchr/testify/require"
	"math/big"
	"mathparse/lexer"
	"testing"
)

func requireNoParseErrors(t *testing.T, p *Parser) {
	require.Equal(t, 0, len(p.errors))
}

// todo: make this pass
func TestParser_ParseIntegerAtom(t *testing.T) {
	input := "2"
	l := lexer.New(input)
	p := New(l)
	expression := p.Parse()
	requireNoParseErrors(t, p)
	require.Equal(t, 1, len(expression.Terms))
	require.Equal(t, big.NewInt(2), expression.Terms[0].Factors[0])
}
