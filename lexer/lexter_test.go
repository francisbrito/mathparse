package lexer

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

type TokenizeTestCase struct {
	input    string
	expected []Token
}

func assertTokenizeEquals(t *testing.T, expected []Token, input string) bool {
	tokens := New(input).Tokenize()
	// NOTE: Skip last token as it is the EndOfFile token.
	return assert.Equal(t, expected, tokens[:len(tokens)-1])
}

func TestTokenizeInteger(t *testing.T) {
	testCases := []TokenizeTestCase{
		{"1", []Token{{Integer, "1"}}},
		{"123", []Token{{Integer, "123"}}},
		{"-1", []Token{{MinusSign, "-"}, {Integer, "1"}}},
		{"-1-", []Token{{Unknown, "-1-"}}},
		{"-1-2", []Token{{Unknown, "-1-2"}}},
		{"-1-2", []Token{{Unknown, "-1-2"}}},
		{"--1", []Token{{Unknown, "--1"}}},
	}
	for _, tc := range testCases {
		assertTokenizeEquals(t, tc.expected, tc.input)
	}
}

func TestTokenizeFloat(t *testing.T) {
	testCases := []TokenizeTestCase{
		{"0.5", []Token{{Float, "0.5"}}},
		{"999.5555", []Token{{Float, "999.5555"}}},
		{"1.23.4", []Token{{Unknown, "1.23.4"}}},
	}
	for _, tc := range testCases {
		if !assertTokenizeEquals(t, tc.expected, tc.input) {
			t.Fatal()
		}
	}
}
