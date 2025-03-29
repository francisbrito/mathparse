package parser

import (
	"fmt"
	"math/big"
	"mathparse/lexer"
	"mathparse/token"
)

type Parser struct {
	l            *lexer.Lexer
	currentToken *token.Token
	peekToken    *token.Token
	errors       []error
}

func New(l *lexer.Lexer) *Parser {
	p := Parser{l: l}
	p.nextToken()
	p.nextToken()
	return &p
}

func (p *Parser) Errors() []error {
	return p.errors
}

func (p *Parser) Parse() *Expression {
	return p.parseExpression()
}

type ParseError struct {
	Message string
	Token   *token.Token
}

func (e *ParseError) String() string {
	return fmt.Sprintf("parse error: %s", e.Message)
}

func (e *ParseError) Error() string {
	return e.String()
}

func (p *Parser) nextToken() {
	p.currentToken = p.peekToken
	pt := p.l.NextToken()
	p.peekToken = &pt
}

func (p *Parser) addParseError(msg string) {
	p.errors = append(p.errors, &ParseError{Message: msg, Token: p.currentToken})
}

func (p *Parser) parseExpression() *Expression {
	var expr Expression
	switch p.currentToken.Type {
	case token.EndOfFile:
		return nil
	case token.Integer:
		expr.LeftTerm = newNumber(p.currentToken.Literal)
		if p.peekToken.Type == token.PlusSign {
			op := Addition
			expr.Op = &op
			p.nextToken()
		} else if p.peekToken.Type == token.MinusSign {
			op := Subtraction
			expr.Op = &op
			p.nextToken()
		} else if p.peekToken.Type == token.Integer {
			p.addParseError("invalid expression")
			return nil
		}
		p.nextToken()
		if p.currentToken.Type != token.EndOfFile {
			expr.RightTerm = p.parseExpression()
		}
	}
	return &expr
}

func newNumber(literal string) *Number {
	value := new(big.Int)
	value.SetString(literal, 0)
	return &Number{Value: value}
}

type Expression struct {
	Op        *TermOperation
	LeftTerm  Atom
	RightTerm Atom
}

type TermOperation int

func (o TermOperation) String() string {
	switch o {
	case Addition:
		return "+"
	case Subtraction:
		return "-"
	default:
		return fmt.Sprintf("unknown operation: %d", o)
	}
}

const (
	Addition TermOperation = iota + 10
	Subtraction
)

type Atom interface {
	Evaluate() *big.Int
	String() string
}

func (e *Expression) Evaluate() *big.Int {
	if e.RightTerm == nil {
		return e.LeftTerm.Evaluate()
	}
	result := new(big.Int)
	left := e.LeftTerm.Evaluate()
	right := e.RightTerm.Evaluate()
	if *e.Op == Addition {
		result.Add(left, right)
	} else if *e.Op == Subtraction {
		result.Sub(left, right)
	}
	return result
}

func (e *Expression) String() string {
	return e.Evaluate().String()
}

type Number struct {
	Value *big.Int
}

func (n *Number) Evaluate() *big.Int {
	return n.Value
}

func (n *Number) String() string {
	return n.Value.String()
}
