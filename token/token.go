package token

type Type string

type Token struct {
	Type    Type
	Literal string
}

const (
	Caret       Type = "CARET"
	EndOfFile        = "EOF"
	Illegal          = "ILLEGAL"
	Integer          = "INTEGER"
	MinusSign        = "MINUS_SIGN"
	PercentSign      = "PERCENT_SIGN"
	PlusSign         = "PLUS_SIGN"
	Asterisk         = "ASTERISK"
	Slash            = "SLASH"
)
