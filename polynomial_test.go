package polynomial

import (
	"math"
	"testing"
)

var first *Polynomial = new(Polynomial)
var second *Polynomial = new(Polynomial)

func SetupTest(t *testing.T) {
	first = new(Polynomial)
	second = new(Polynomial)
	t.Log("first")
	first.AddTerm(1, 2.1)
	first.AddTerm(4, 2)
	first.AddTerm(3, 1)
	first.AddTerm(0, 1.3)
	first.AddTerm(4, 0.3)
	t.Log(first)

	t.Log("second")
	second.AddTerm(4, -2.3)
	second.AddTerm(2, 1)
	second.AddTerm(0, -1.3)
	second.AddTerm(1, 0.3)
	t.Log(second)
}

func TestPolynomial_Add(t *testing.T) {
	SetupTest(t)
	t.Log("add first and second")
	third := first.Add(second)
	t.Log(first)
	t.Log(second)
	t.Log(third)

	expected := [3]float64{2.4, 1, 1}
	actual := []float64{}
	for tmp := third.head; tmp != nil; tmp = tmp.next {
		actual = append(actual, tmp.coeff)
	}
	if len(actual) != len(expected) {
		t.Error("addition failed")
		return
	}
	for i := range expected {
		if math.Abs(actual[i]-expected[i]) > Tolerance {
			t.Error("addition failed")
		}
	}
}

func TestPolynomial_Multiply(t *testing.T) {
	SetupTest(t)
	t.Log("multiply first by 0.2")
	third := first.Multiply(0.2)
	t.Log(first)
	t.Log(third)

	expected := [4]float64{0.26, 0.42, 0.2, 0.46}
	actual := []float64{}
	for tmp := third.head; tmp != nil; tmp = tmp.next {
		actual = append(actual, tmp.coeff)
	}
	if len(actual) != len(expected) {
		t.Error("multiplication failed")
		return
	}

	for i := range expected {
		if math.Abs(actual[i]-expected[i]) > Tolerance {
			t.Error("multiplication failed")
		}
	}
}

func TestPolynomial_Diff(t *testing.T) {
	SetupTest(t)
	t.Log("differentiate first")
	third := first.Diff()
	t.Log(first)
	t.Log(third)

	expected := [3]float64{2.1, 3, 9.2}
	actual := []float64{}
	for tmp := third.head; tmp != nil; tmp = tmp.next {
		actual = append(actual, tmp.coeff)
	}
	if len(actual) != len(expected) {
		t.Error("differentiation failed")
		return
	}

	for i := range expected {
		if math.Abs(actual[i]-expected[i]) > Tolerance {
			t.Error("differentiation failed")
		}
	}
}

func TestPolynomial_Evaluate(t *testing.T) {
	SetupTest(t)
	t.Log("eval first at x = 1.5")
	actual := first.Evaluate(1.5)
	t.Log(first)
	t.Log(actual)

	expected := 19.46875

	if math.Abs(actual-expected) > Tolerance {
		t.Error("evaluation failed")
	}
}
