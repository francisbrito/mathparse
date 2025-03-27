package ast

import (
	"github.com/stretchr/testify/require"
	"math/big"
	"testing"
)

func TestNumber_Evaluate(t *testing.T) {
	n := newNumber(5)
	require.Equal(t, big.NewFloat(5), n.Evaluate())
}

func TestNumber_String(t *testing.T) {
	n := newNumber(5)
	require.Equal(t, big.NewFloat(5).SetPrec(4).String(), n.String())
}

func TestFactor_Evaluate(t *testing.T) {
	f := Factor{Left: newNumber(5)}
	require.Equal(t, big.NewFloat(5), f.Evaluate())
	op := Multiplication
	f = Factor{Left: newNumber(5), FactorOp: &op, Right: newNumber(5)}
	require.Equal(t, 0, f.Evaluate().Cmp(big.NewFloat(25)))
}

func TestFactor_String(t *testing.T) {
	f := Factor{Left: newNumber(5)}
	require.Equal(t, "5", f.String())
	f = newFactor(5, Multiplication, 5)
	require.Equal(t, "5 * 5", f.String())

}

func TestTerm_Evaluate(t *testing.T) {
	term := Term{Left: &Factor{Left: newNumber(5)}}
	require.Equal(t, big.NewFloat(5), term.Evaluate())
	term = newTerm(5, Addition, 5)
	require.Equal(t, 0, term.Evaluate().Cmp(big.NewFloat(10)))
}

func TestTerm_String(t *testing.T) {
	term := Term{Left: &Factor{Left: newNumber(5)}}
	require.Equal(t, "5", term.String())
	term = newTerm(5, Addition, 5)
	require.Equal(t, "5 + 5", term.String())
}

func newNumber(n float64) Number {
	return Number{Float: big.NewFloat(n)}
}

func newFactor(l float64, op FactorOperation, r float64) Factor {
	return Factor{Left: newNumber(l), FactorOp: &op, Right: newNumber(r)}
}

func newTerm(l float64, op TermOperation, r float64) Term {
	return Term{
		TermOp: &op,
		Left:   &Factor{Left: newNumber(l)},
		Right:  &Factor{Left: newNumber(r)},
	}
}
