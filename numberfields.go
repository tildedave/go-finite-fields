package main

type NumberFieldSignature struct {
	numReal    int
	numComplex int
}

func signLeadingCoefficient(p []int64) int {
	if PolynomialLeadingCoefficient(p) < 0 {
		return -1
	}

	return 1
}

// ComputeNumberFieldSignature will the signature r1, r2 of the passed in polynomial using Sturm's Theorem. r1 + 2*r2 = n
// A Course In Computational Number Theory, pg 156
func ComputeNumberFieldSignature(tPoly []int64) NumberFieldSignature {
	result := NumberFieldSignature{}
	if PolynomialDegree(tPoly) == 0 {
		return result
	}

	a := PolynomialPrimitivePart(tPoly)
	b := PolynomialPrimitivePart(PolynomialDerivative(tPoly, 0))

	g := int64(1)
	h := int64(1)
	s := signLeadingCoefficient(a)
	n := PolynomialDegree(a)
	var t int
	if (n-1)%2 == 0 {
		t = s
	} else {
		t = -s
	}
	result.numReal = 1

	for {
		// Step 2
		delta := PolynomialDegree(a) - PolynomialDegree(b)
		_, r := PolynomialPsuedoDivide(a, b)

		if len(r) == 0 {
			panic("T was not squarefree")
		}
		if PolynomialLeadingCoefficient(b) > 0 || delta%2 == 1 {
			r = PolynomialScalar(r, -1, 0)
		}

		// Step 3
		if signLeadingCoefficient(r) != s {
			s = -s
			result.numReal--
		}

		deg := PolynomialDegree(r)
		var expectedSign int
		if deg%2 == 0 {
			expectedSign = t
		} else {
			expectedSign = -t
		}

		if signLeadingCoefficient(r) != expectedSign {
			t = -t
			result.numReal++
		}

		// Step 4
		if PolynomialDegree(r) == 0 {
			if (n-result.numReal)%2 == 1 {
				panic("Bug in algorithm")
			}

			result.numComplex = (n - result.numReal) / 2
			return result
		}

		a = make([]int64, len(b))
		copy(a, b)
		b = make([]int64, len(r))
		for i := range r {
			b[i] = r[i] / (g * IntExp(h, int64(delta)))
		}
		leadCoeff := PolynomialLeadingCoefficient(a)
		if leadCoeff < 0 {
			g = -leadCoeff
		} else {
			g = leadCoeff
		}
		// h = h^(1 - delta)g^delta
		if delta != 0 {
			h = g
		}
	}
}
