package ast

import (
	"fmt"
	"mathparse/token"
)

type Expression interface {
	expression()
	String() string
	Literal() string
}

type IntegerExpression struct {
	Token token.Token
	Value int64
}

func (i *IntegerExpression) expression() {}

func (i *IntegerExpression) Literal() string {
	return i.Token.Literal
}

func (i *IntegerExpression) String() string {
	return fmt.Sprintf("%s: %d", i.Token.Type, i.Value)
}
