package parser

import (
	"github.com/stretchr/testify/assert"
	"mathparse/ast"
	"mathparse/lexer"
	"testing"
)

func TestParseInteger(t *testing.T) {
	input := "5"
	l := lexer.New(input)
	p := New(l)
	expressions := p.Parse()
	if !assert.Equal(t, 1, len(expressions)) {
		t.Fatal()
	}
	expr, ok := expressions[0].(*ast.IntegerExpression)
	if !ok {
		t.Fatal("not an IntegerExpression")
	}
	if !assert.Equal(t, int64(5), expr.Value) {
		t.Fatal()
	}
}

func TestParseFloat(t *testing.T) {
	input := "1.5"
	l := lexer.New(input)
	p := New(l)
	expressions := p.Parse()
	if !assert.Equal(t, 1, len(expressions)) {
		t.Fatal()
	}
	expr, ok := expressions[0].(*ast.FloatExpression)
	if !ok {
		t.Fatal("not a FloatExpression")
	}
	if !assert.Equal(t, 1.5, expr.Value) {
		t.Fatal()
	}
}
