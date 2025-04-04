package lexer

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"mathparse/token"
	"testing"
)

type testCase struct {
	input    string
	expected []token.Token
}

var digitTokens []token.Token
var arithmeticOpTokens []token.Token

func init() {
	for _, digit := range "1234567890" {
		literal := string(digit)
		digitTokens = append(digitTokens, token.Token{Type: token.Integer, Literal: literal})
	}
	for _, op := range "+-*/%^" {
		literal := string(op)
		var tokenType token.Type
		switch op {
		case '+':
			tokenType = token.PlusSign
		case '-':
			tokenType = token.MinusSign
		case '*':
			tokenType = token.Asterisk
		case '/':
			tokenType = token.Slash
		case '%':
			tokenType = token.PercentSign
		case '^':
			tokenType = token.Caret
		default:
			panic(fmt.Errorf("cannot setup test case, invalid operator: %q", op))
		}
		arithmeticOpTokens = append(arithmeticOpTokens, token.Token{Type: tokenType, Literal: literal})
	}
}

func newTestCase(input string, tokens ...token.Token) *testCase {
	return &testCase{input: input, expected: tokens}
}

func TestLexer_NextToken(t *testing.T) {
	var lexer *Lexer
	testCases := []*testCase{
		// can process empty input
		newTestCase("", token.Token{Type: token.EndOfFile}),
		// can process illegal input
		newTestCase("ILLEGAL", token.Token{Type: token.Illegal, Literal: "ILLEGAL"}, token.Token{Type: token.EndOfFile, Literal: ""}),
		// can skip whitespace
		newTestCase(" ", token.Token{Type: token.EndOfFile}),
		newTestCase("		", token.Token{Type: token.EndOfFile}),
		// can tokenize integers
		newTestCase("1234567890", token.Token{Type: token.Integer, Literal: "1234567890"}),
		// can tokenize floating point numbers
		newTestCase("1.2", token.Token{Type: token.Float, Literal: "1.2"}),
		newTestCase("1.2.3", token.Token{Type: token.Illegal, Literal: "1.2.3"}),
		// can tokenize parentheses
		newTestCase("(", token.Token{Type: token.OpeningParentheses, Literal: "("}),
		newTestCase(")", token.Token{Type: token.ClosingParentheses, Literal: ")"}),
		// can tokenize a complex expressions
		newComplexTestCase(),
	}
	// can tokenize every digit
	for _, digitToken := range digitTokens {
		testCases = append(testCases, newTestCase(digitToken.Literal, digitToken))
	}
	// can tokenize arithmetic operators
	for _, opToken := range arithmeticOpTokens {
		testCases = append(testCases, newTestCase(opToken.Literal, opToken))
	}
	for _, tc := range testCases {
		lexer = New(tc.input)
		for _, expectedToken := range tc.expected {
			actualToken := lexer.NextToken()
			if !assert.Equal(t, expectedToken, actualToken, "input: %q", tc.input) {
				t.Fatal()
			}
		}
	}
}

func newComplexTestCase() *testCase {
	return newTestCase("1 + 2 - 3 * 4 / 5 ^ (6 % 7)",
		token.Token{
			Type:    token.Integer,
			Literal: "1",
		},
		token.Token{
			Type:    token.PlusSign,
			Literal: "+",
		},
		token.Token{
			Type:    token.Integer,
			Literal: "2",
		},
		token.Token{
			Type:    token.MinusSign,
			Literal: "-",
		},
		token.Token{
			Type:    token.Integer,
			Literal: "3",
		},
		token.Token{
			Type:    token.Asterisk,
			Literal: "*",
		},
		token.Token{
			Type:    token.Integer,
			Literal: "4",
		},
		token.Token{
			Type:    token.Slash,
			Literal: "/",
		},
		token.Token{
			Type:    token.Integer,
			Literal: "5",
		},
		token.Token{
			Type:    token.Caret,
			Literal: "^",
		},
		token.Token{
			Type:    token.OpeningParentheses,
			Literal: "(",
		},
		token.Token{
			Type:    token.Integer,
			Literal: "6",
		},
		token.Token{
			Type:    token.PercentSign,
			Literal: "%",
		},
		token.Token{
			Type:    token.Integer,
			Literal: "7",
		},
		token.Token{
			Type:    token.ClosingParentheses,
			Literal: ")",
		},
	)
}
