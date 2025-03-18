package lexer

import (
	"github.com/stretchr/testify/assert"
	"mathparse/token"
	"testing"
)

var digitTokens []token.Token

func init() {
	for _, digit := range "1234567890" {
		literal := string(digit)
		digitTokens = append(digitTokens, token.Token{Type: token.Integer, Literal: literal})
	}
}

type testCase struct {
	input    string
	expected []token.Token
}

func newTestCase(input string, tokens ...token.Token) *testCase {
	return &testCase{input: input, expected: tokens}
}

func TestLexer_NextToken(t *testing.T) {
	var lexer *Lexer
	testCases := []*testCase{
		// can process empty input
		newTestCase("", token.Token{Type: token.EndOfFile}),
		// can tokenize illegal tokens
		newTestCase("ILLEGAL", token.Token{Type: token.Illegal, Literal: "ILLEGAL"}),
		// can skip whitespace
		newTestCase(" ", token.Token{Type: token.EndOfFile}),
		newTestCase("		", token.Token{Type: token.EndOfFile}),
		// can tokenize integers
		newTestCase("1234567890", token.Token{Type: token.Integer, Literal: "1234567890"}),
	}
	// can tokenize every digit
	for _, digitToken := range digitTokens {
		testCases = append(testCases, newTestCase(digitToken.Literal, digitToken))
	}
	for _, tc := range testCases {
		lexer = New(tc.input)
		for _, expectedToken := range tc.expected {
			actualToken := lexer.NextToken()
			if !assert.Equal(t, expectedToken, actualToken) {
				t.Fatal()
			}
		}
	}
}
