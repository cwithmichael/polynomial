# polynomial

[![Go Report Card](https://goreportcard.com/badge/github.com/cwithmichael/polynomial)](https://goreportcard.com/report/github.com/cwithmichael/polynomial)

implements routines for manipulating univariate polynomials over reals

This is a rewrite of the Java code from [here](https://www.cs.cmu.edu/~adamchik/15-121/lectures/Linked%20Lists/code/Polynomial.java) in Go

## Example Usage
```
package main
import (
	"github.com/cwithmichael/polynomial"
	"fmt"
)

func main() {
	p := new(polynomial.Polynomial)
	p2 := new(polynomial.Polynomial)
	p.AddTerm(1, 1)
	p.AddTerm(2, 2)
	p2.AddTerm(3, 1)
	p2.AddTerm(4, 2)
	sum, _ := p.Add(p2)
	fmt.Printf("(%v) + (%v) = %v\n", p, p2, sum)

	p3 := new(polynomial.Polynomial)
	p3.AddTerm(3.12, 2)
	p3.AddTerm(4, 6)
	p3.AddTerm(-5, 4)
	diff, _ := p3.Diff()
	fmt.Printf("d/dx %v = %v\n", p3, diff)

	res := diff.Evaluate(20)
	fmt.Printf("f = %v\n", diff)
	fmt.Printf("f(20) = %v\n", res)
}

```

```
go run main.go
(1.00x + 2.00x^2) + (3.00x + 4.00x^2) = 4.00x + 6.00x^2
d/dx 3.12x^2 - 5.00x^4 + 4.00x^6 = 6.24x - 20.00x^3 + 24.00x^5
f = 6.24x - 20.00x^3 + 24.00x^5
f(20) = 7.66401248e+07
```
