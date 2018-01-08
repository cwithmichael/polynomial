package polynomial

import (
	"math"
	"strconv"
)

type Monomial struct {
	exp int64
	coeff float64
	next *Monomial
}

func (m *Monomial) String() string {
	form := strconv.FormatFloat(math.Abs(m.coeff), 'f', 2, 64)

	switch m.exp {
	case 0:
		return form
	case 1:
		return form + "*x"
	default:
		return form + "*x^" + strconv.FormatInt(m.exp, 10)
	}
}

func (m *Monomial) Equals(mono *Monomial) bool {
	return m.exp == mono.exp && m.coeff == mono.coeff;
}