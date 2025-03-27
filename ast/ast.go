package ast

import (
	"fmt"
	"math/big"
)

// precision, in bits
const operationPrecision = 100

type TermOperation int

func (o TermOperation) String() string {
	switch o {
	case Addition:
		return "+"
	case Subtraction:
		return "-"
	}
	panic("invalid operation")
}

const (
	Addition TermOperation = iota
	Subtraction
)

type FactorOperation int

func (o FactorOperation) String() string {
	switch o {
	case Multiplication:
		return "*"
	case Division:
		return "/"
	}
	panic("invalid operation")
}

const (
	Multiplication FactorOperation = iota
	Division
)

type Atom interface {
	Evaluate() *big.Float
	String() string
}

type Expression struct {
	Terms []Term
}

type Term struct {
	TermOp      *TermOperation
	Left, Right Atom
}

func (t *Term) Evaluate() *big.Float {
	result := new(big.Float).SetPrec(operationPrecision)
	if t.TermOp == nil || t.Right == nil {
		return t.Left.Evaluate()
	}
	left, right := t.Left.Evaluate(), t.Right.Evaluate()
	switch *t.TermOp {
	case Addition:
		result.Add(left, right)
	case Subtraction:
		result.Add(left, right)
	default:
		panic("invalid operation")
	}
	return result
}

func (t *Term) String() string {
	if t.TermOp == nil || t.Right == nil {
		return t.Left.String()
	}
	return fmt.Sprintf("%s %s %s", t.Left.String(), t.TermOp.String(), t.Right.String())
}

type Factor struct {
	FactorOp    *FactorOperation
	Left, Right Atom
}

func (f *Factor) Evaluate() *big.Float {
	result := new(big.Float).SetPrec(operationPrecision)
	if f.FactorOp == nil || f.Right == nil {
		return f.Left.Evaluate()
	}
	left, right := f.Left.Evaluate(), f.Right.Evaluate()
	switch *f.FactorOp {
	case Multiplication:
		result.Mul(left, right)
	case Division:
		result.Quo(left, right)
	default:
		panic("invalid operation")
	}
	return result
}

func (f *Factor) String() string {
	if f.FactorOp == nil || f.Right == nil {
		return f.Left.String()
	}
	return fmt.Sprintf("%s %s %s", f.Left.String(), f.FactorOp.String(), f.Right.String())
}

type Number struct {
	*big.Float
}

func (n Number) Evaluate() *big.Float {
	return n.Float
}

func (n Number) String() string {
	return n.Float.String()
}

func NewNumber(value *big.Float) Number {
	return Number{value}
}
