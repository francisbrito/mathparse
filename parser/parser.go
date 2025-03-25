package parser

import (
	"fmt"
	"mathparse/ast"
	"mathparse/lexer"
	"mathparse/token"
)

type Parser struct {
	l            *lexer.Lexer
	currentToken token.Token
	peekToken    token.Token
	errors       []Error
}

type Error struct {
	Token   token.Token
	Message string
}

func (e Error) String() string {
	return fmt.Sprintf("parse error: %s: %s", e.Token, e.Message)
}

func (e Error) Error() string {
	return e.String()
}

func New(l *lexer.Lexer) *Parser {
	p := &Parser{l: l}
	p.nextToken()
	p.nextToken()
	return p
}

func (p *Parser) nextToken() {
	p.currentToken = p.peekToken
	p.peekToken = p.l.NextToken()
}

func (p *Parser) Parse() ast.Expression {
	var expression ast.Expression
	for p.currentToken.Type != token.EndOfFile {
		switch p.currentToken.Type {
		case token.Illegal:
		default:
			p.addError("illegal token")
		}
		p.nextToken()
	}
	return expression
}

func (p *Parser) addError(message string) {
	p.errors = append(p.errors, Error{Token: p.currentToken, Message: message})
}
