package token

type Type string

type Token struct {
	Type    Type
	Literal string
}

const (
	EndOfFile Type = "EOF"
	Illegal        = "ILLEGAL"
	Integer        = "INTEGER"
)
