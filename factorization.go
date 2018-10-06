package main

import (
	"math/rand"
)

func createMatrix(size int) [][]int64 {
	matrix := make([][]int64, size)
	for i := range matrix {
		matrix[i] = make([]int64, size)
	}

	return matrix
}

func computeNullSpace(matrix [][]int64, n int, char int64) ([][]int64, int) {
	cols := make([]int, n)
	for i := range cols {
		cols[i] = -1
	}
	r := 0

	vecs := make([][]int64, n)

OUTER:
	for k := 0; k < n; k++ {
		for j := 0; j < n; j++ {
			if matrix[k][j] != 0 && cols[j] < 0 {
				// time to compute
				// multiply column j of matrix by -1/matrix[k][j]
				a := matrix[k][j]
				// -1/a
				adj := ((char - 1) * ModInverse(a, char)) % char

				for y := n - 1; y >= k; y-- {
					matrix[y][j] = (matrix[y][j] * adj) % char
				}

				// TOOD(perf): Possibly defer modulus until end of computation
				for y := n - 1; y >= k; y-- {
					for x := 0; x < n; x++ {
						if x == j {
							continue
						} else {
							matrix[y][x] = (matrix[y][x] + matrix[k][x]*matrix[y][j]) % char
						}
					}
				}

				cols[j] = k

				continue OUTER
			}

		}

		// output vector
		vec := make([]int64, n)
	VEC:
		for j := 0; j < n; j++ {
			for s := range cols {
				if cols[s] == j {
					vec[j] = matrix[k][s]
					continue VEC
				}
			}

			if j == k {
				vec[j] = 1
			} else {
				vec[j] = 0
			}
		}

		vecs[r] = PolynomialTrunc(vec)
		r++
	}

	return vecs[0:r], r

}

// FactorBerlekamp factors the given squarefree polynomial f into irreducible factors using
// Berlekamp's algorithm.  This follows the implementation given in TAoCP 6.4.2.
func FactorBerlekamp(f []int64, char int64) [][]int64 {
	n := len(f) - 1
	matrix := createMatrix(n)
	unit := []int64{0, 1}

	// Compute Q - I where I is the unit matrix
	// The kth row is x^(kp) mod f.  We skip the first row because Q - I is always 0.
	for k := 1; k < n; k++ {
		t := make([]int64, n)
		p := PolynomialModExp(unit, int64(k)*char, f, char)
		copy(t, p)
		matrix[k] = t
		matrix[k][k] = (matrix[k][k] + char - int64(1)) % char
	}

	// Find null space for Q - I
	vecs, numSolutions := computeNullSpace(matrix, n, char)

	// Polynomial is irreducible
	if numSolutions == 1 {
		return [][]int64{}
	}

	// For each vector, compute GCD for u(x), vec - s for 0 <= s < p.
	// The result gives a nontrivial factorization of u.
	// We push each nontrivial factorization of u onto a list and
	// then continue to reduce.
	factors := make([][]int64, numSolutions)
	foundFactors := 0

	isFactorNew := func(p []int64) bool {
		for i := 0; i < foundFactors; i++ {
			f := factors[i]
			if PolynomialsAreEqual(p, f) {
				// already found this factor, nothing to do
				return false
			}
		}

		return true
	}

	// vecs[1] provides a non-trivial factoring of f
	vec := vecs[1]
	for s := int64(0); s < char; s++ {
		unit := []int64{char - s}
		p1 := PolynomialAdd(vec, unit, char)
		p := PolynomialGcd(p1, f, char)
		if len(p) > 1 {
			// yay!
			factors[foundFactors] = p
			foundFactors++
		}
	}

	// If vecs[1] split f into irreducibles we're done
	if foundFactors == numSolutions {
		return factors
	}

	// Otherwise, we use the other factorings to narrow down
	for _, vec := range vecs[2:] {
		for i := 0; i < foundFactors; i++ {
			factor := factors[i]
			for s := int64(0); s < char; s++ {
				unit := []int64{char - s}
				p1 := PolynomialAdd(vec, unit, char)
				p := PolynomialGcd(p1, factor, char)

				if len(p) > 1 {
					if isFactorNew(p) {
						// GCD(p, factor) was non-trivial, so factor was not irreducible
						factors[i] = p

						// See if we've seen factor / q yet
						q, _ := PolynomialDivide(p, factor, char)
						if isFactorNew(q) {
							factors[foundFactors] = q
							foundFactors++
							if foundFactors == numSolutions {
								return factors
							}
						}
					}
				}
			}
		}
	}

	panic("Bug in factorization, should never get here")
}

type DistinctDegreeFactor struct {
	factor []int64
	degree int
}

// FactorDistinctDegree uses the distinct degree factorization method from TAoCP 4.6.2.
// It assumes f is squarefree.
// The basic idea is to use GCDs with x^p^d - x to find irreducibles of degree d that divide f.
// The end result is a non-trivial factorization of f.  The individual polynomials may not be irreducible.
func FactorDistinctDegree(f []int64, char int64) []DistinctDegreeFactor {
	v := f
	unit := []int64{0, 1}
	w := unit
	d := 0
	solutions := make([]DistinctDegreeFactor, PolynomialDegree(f))
	numSolutions := 0

	// invariant: w = x^p^d mod v

	for {
		if d+1 > PolynomialDegree(v)/2 {
			if len(v) > 1 {
				// TODO: not sure if degree is correct here
				solutions[numSolutions] = DistinctDegreeFactor{
					degree: PolynomialDegree(v),
					factor: v}
				numSolutions++
			}
			return solutions[0:numSolutions]
		}

		d = d + 1
		w = PolynomialModExp(w, char, v, char)
		// g_d(x) = gcd(w - x, v(x))
		g := PolynomialTrunc(PolynomialSubtract(w, unit, char))

		if len(g) == 0 {
			// x^p^d - x cleanly divides v.  this means that v is a product of
			// all irreducibles of degree d.  we don't need to factor it further.

			solutions[numSolutions] = DistinctDegreeFactor{degree: d, factor: v}
			numSolutions++

			return solutions[0:numSolutions]
		}

		gd := PolynomialGcd(g, v, char)

		if len(gd) > 1 {
			v, _ = PolynomialDivide(gd, v, char)
			w = PolynomialMod(w, v, char)

			// NOTE: if degree(gd) > d, gd is not irreducible
			solutions[numSolutions] = DistinctDegreeFactor{degree: d, factor: gd}
			numSolutions++
		}
	}
}

func FactorCantorZassenhaus(r *rand.Rand, f []int64, char int64) [][]int64 {
	if char == 2 {
		panic("Don't yet support factorization for p = 2")
	}

	factors := make([][]int64, len(f))
	numFactors := 0
	distinctDegreeFactors := FactorDistinctDegree(f, char)

	for _, solution := range distinctDegreeFactors {
		degreeFactors := make([][]int64, 0)

		// split solution.factor into its distinct factors, all with the same degree
		queue := make([][]int64, 1)
		queue[0] = solution.factor

		for len(queue) > 0 {
			first := queue[0]
			queue = queue[1:]

			if PolynomialDegree(first) == solution.degree {
				degreeFactors = append(degreeFactors, first)
			} else {
				g := FactorEqualDegree(r, first, solution.degree, char)
				q, _ := PolynomialDivide(g, first, char)
				queue = append(append(queue, q), g)
			}
		}

		for _, factor := range degreeFactors {
			factors[numFactors] = factor
			numFactors++
		}
	}

	return factors[0:numFactors]
}

// FactorEqualDegree uses a random generator to find an irreducible factor of f.
func FactorEqualDegree(r *rand.Rand, f []int64, d int, char int64) []int64 {
	const numIterations = 500

	for i := 0; i < numIterations; i++ {
		factor := factorEqualDegreeAttempt(r, f, d, char)
		if factor != nil {
			return factor
		}
	}

	panic("No factor after trying a bunch")
}

func factorEqualDegreeAttempt(r *rand.Rand, f []int64, d int, char int64) []int64 {
	a := PolynomialRandom(r, PolynomialDegree(f)-1, char)
	p := PolynomialGcd(f, a, char)

	if len(p) > 1 {
		// somehow p is a factor
		return p
	}

	p = PolynomialModExp(a, Exp(char, d)-1, f, char)
	p[0] = (p[0] + 1) % char
	g := PolynomialGcd(p, f, char)

	if len(g) > 1 {
		return g
	}

	return nil
}
