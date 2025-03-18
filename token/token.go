package token

type Type string

type Token struct {
	Type    Type
	Literal string
}

const (
	Asterisk           Type = "ASTERISK"
	Caret                   = "CARET"
	ClosingParentheses      = "CLOSING_PARENTHESES"
	EndOfFile               = "EOF"
	Float                   = "FLOAT"
	Illegal                 = "ILLEGAL"
	Integer                 = "INTEGER"
	MinusSign               = "MINUS_SIGN"
	OpeningParentheses      = "OPENING_PARENTHESES"
	PercentSign             = "PERCENT_SIGN"
	PlusSign                = "PLUS_SIGN"
	Slash                   = "SLASH"
)
