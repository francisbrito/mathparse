package parser

import (
	"fmt"
	"mathparse/ast"
	"mathparse/lexer"
	"mathparse/token"
	"strconv"
)

type Parser struct {
	l            *lexer.Lexer
	currentToken token.Token
	peekToken    token.Token
}

func New(l *lexer.Lexer) *Parser {
	p := &Parser{l: l}
	// populate current and peek token fields
	p.nextToken()
	p.nextToken()
	return p
}

func (p *Parser) Parse() []ast.Expression {
	var expressions []ast.Expression
	for p.currentToken.Type != token.EndOfFile {
		switch p.currentToken.Type {
		case token.Integer:
			expressions = append(expressions, p.parseInteger())
		case token.Float:
			expressions = append(expressions, p.parseFloat())
		default:
			// todo: implement this
			panic("invalid expression")
		}
		p.nextToken()
	}
	return expressions
}

func (p *Parser) nextToken() {
	p.currentToken = p.peekToken
	p.peekToken = p.l.NextToken()
}

func (p *Parser) parseInteger() *ast.IntegerExpression {
	value, err := strconv.ParseInt(p.currentToken.Literal, 0, 64)
	if err != nil {
		// todo: implement this
		panic(fmt.Errorf("parse error: unable to parse integer: %q", p.currentToken.Literal))
	}
	return &ast.IntegerExpression{Value: value, Token: p.currentToken}
}

func (p *Parser) parseFloat() *ast.FloatExpression {
	value, err := strconv.ParseFloat(p.currentToken.Literal, 0)
	if err != nil {
		// todo: implement this
		panic(fmt.Errorf("parse error: unable to parse float: %q", p.currentToken.Literal))
	}
	return &ast.FloatExpression{Value: value, Token: p.currentToken}
}
