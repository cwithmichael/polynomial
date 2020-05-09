// Package polynomial implements routines for
// manipulating univariate polynomials over reals.
package polynomial

import (
	"bytes"
	"fmt"
	"math"
)

const (
	// Tolerance of error
	Tolerance = 0.00000001
)

// Polynomial represents a linked list of Monomials
type Polynomial struct {
	head *Monomial
}

// AddTerm adds a new term into the polynomial, assuming that the polynomial
// is sorted in order from smallest to largest exponent.
func (p *Polynomial) AddTerm(coeff float64, exp int64) error {
	if exp < 0 {
		return fmt.Errorf("math: exponent cannot be a negative number %d", exp)
	}
	if math.Abs(coeff) < Tolerance {
		return fmt.Errorf("math: coefficient %g is less than tolerance %g", coeff, Tolerance)
	}
	if p.head == nil || exp < p.head.exp {
		p.head = &Monomial{exp, coeff, p.head}
		return nil
	}

	cur := p.head
	var prev *Monomial

	for cur != nil && exp > cur.exp {
		prev = cur
		cur = cur.next
	}
	if cur == nil || exp != cur.exp {
		prev.next = &Monomial{exp, coeff, cur}
	} else {
		cur.coeff += coeff
		if math.Abs(cur.coeff) < Tolerance {
			if prev != nil {
				prev.next = cur.next
			} else {
				p.head = p.head.next
			}
		}
	}
	return nil
}

// String returns the polynomial as a string
func (p *Polynomial) String() string {
	var buffer bytes.Buffer
	for count, tmp := 0, p.head; tmp != nil; count, tmp = count+1, tmp.next {
		if count == 0 {
			buffer.WriteString(tmp.String())
		} else {
			if tmp.coeff < 0 {
				buffer.WriteString(" - " + tmp.String())
			} else {
				buffer.WriteString(" + " + tmp.String())
			}
		}
	}

	return buffer.String()
}

// Add adds two polynomials
// The method does not change the original polynomial.
func (p *Polynomial) Add(poly *Polynomial) (*Polynomial, error) {
	res, err := p.clone()

	if err != nil {
		return nil, err
	}

	for tmp := poly.head; tmp != nil; tmp = tmp.next {
		err := res.AddTerm(tmp.coeff, tmp.exp)
		if err != nil {
			return nil, err
		}
	}

	return res, nil
}

func (p *Polynomial) clone() (*Polynomial, error) {
	res := new(Polynomial)

	for tmp := p.head; tmp != nil; tmp = tmp.next {
		err := res.AddTerm(tmp.coeff, tmp.exp)
		if err != nil {
			return nil, err
		}
	}

	return res, nil
}

func (p *Polynomial) equals(poly *Polynomial) bool {
	tmp1 := p.head
	tmp2 := poly.head

	for tmp1 != nil && tmp2 != nil {
		if !tmp1.Equals(tmp2) {
			return false
		}
		tmp1 = tmp1.next
		tmp2 = tmp2.next
	}
	return true
}

// Multiply multiplies by a number
// The method does not change the original polynomial.
func (p *Polynomial) Multiply(num float64) (*Polynomial, error) {
	res, err := p.clone()

	if err != nil {
		return nil, err
	}
	for tmp := res.head; tmp != nil; tmp = tmp.next {
		tmp.coeff *= num
	}
	return res, nil
}

// Diff returns a new polynomial that is the derivative of this polynomial.
func (p *Polynomial) Diff() (*Polynomial, error) {
	res := new(Polynomial)
	for tmp := p.head; tmp != nil; tmp = tmp.next {
		if tmp.exp != 0 {
			err := res.AddTerm(tmp.coeff*float64(tmp.exp), tmp.exp-1)
			if err != nil {
				return nil, err
			}
		}
	}

	return res, nil
}

// Evaluate computes the polynomial at x = value
func (p *Polynomial) Evaluate(value float64) float64 {

	res := 0.0

	for tmp := p.head; tmp != nil; tmp = tmp.next {
		res += tmp.coeff * math.Pow(value, float64(tmp.exp))
	}

	return res
}
