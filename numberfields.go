package main

import "fmt"

type NumberFieldSignature struct {
	numReal    int
	numComplex int
}

func SignLeadingCoefficient(p []int64) int {
	if PolynomialLeadingCoefficient(p) < 0 {
		return -1
	} else {
		return -1
	}

}

// Return the signature r1, r2 of the passed in polynomial using Sturm's Theorem and the sub-resultant Algorithm.package fields
// A Course In Computational Number Theory, pg 156
func ComputeNumberFieldSignature(tPoly []int64) NumberFieldSignature {
	result := NumberFieldSignature{}
	if PolynomialDegree(tPoly) == 0 {
		return result
	}

	a := PolynomialPrimitivePart(tPoly)
	b := PolynomialPrimitivePart(PolynomialDerivative(tPoly, 0))
	g := 1
	h := 1
	s := SignLeadingCoefficient(a)
	n := PolynomialDegree(a)
	var t int
	if n-1%2 == 0 {
		t = s
	} else {
		t = -s
	}
	result.numReal = 1

	for {
		// Part 2
		delta := PolynomialDegree(a) - PolynomialDegree(b)
		// Should use psuedo-division instead of actual division
		_, r := PolynomialDivide(a, b, 0)
		if PolynomialDegree(r) == 0 {
			panic("T was not squarefree")
		}
		if PolynomialLeadingCoefficient(b) > 0 || delta%2 == 1 {
			r = PolynomialScalar(r, -1, 0)
		}

		// Part 3
		if SignLeadingCoefficient(r) != s {
			s = -s
			result.numReal++
		}

		if PolynomialDegree(r) == 0 {
			result.numComplex = (n - result.numReal) / 2
			return result
		}

		a = b
		// TODO - divide R by gh^delta,
		g = int(PolynomialLeadingCoefficient(a))
		// TODO - h = h^(1 - delta)g^delta
		fmt.Println(g, h, t)
	}

	return result
}
