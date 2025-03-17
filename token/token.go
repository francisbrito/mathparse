package token

type Type string

type Token struct {
	Type    Type
	Literal string
}

const (
	Integer   Type = "INTEGER"
	Float          = "FLOAT"
	Unknown        = "UNKNOWN"
	EndOfFile      = "EOF"
	MinusSign      = "MINUS_SIGN"
)
