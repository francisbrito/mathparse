package parser

import (
	"fmt"
	"math/big"
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

func (p *Parser) Errors() []string {
	var errorMessages []string
	for _, e := range p.errors {
		errorMessages = append(errorMessages, e.Error())
	}
	return errorMessages
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
		case token.Integer:
			expression.Terms = append(expression.Terms, p.parseInteger())
		case token.Illegal:
		default:
			p.addError("illegal token")
			break
		}
		p.nextToken()
	}
	if len(p.errors) > 0 {
		expression.Terms = make([]ast.Term, 0)
	}
	return expression
}

func (p *Parser) addError(message string) {
	p.errors = append(p.errors, Error{Token: p.currentToken, Message: message})
}

func (p *Parser) parseInteger() ast.Term {
	value := new(big.Float)
	value.SetPrec(100)
	value.SetString(p.currentToken.Literal)
	return ast.Term{Left: ast.NewNumber(value)}
}
