package ast

type Expression struct {
	Terms []Term
}

type Term struct {
	Factors []Atom
}

// Atom can be a number (integer, float) or an expression.
type Atom interface{}
