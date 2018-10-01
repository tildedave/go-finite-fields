package main

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

// FactorBerlekamp factors the given polynomial f into irreducible factors using Berlekamp's
// algorithm.  This follows the implementation given in TAoCP 6.4.2.
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

	// For each vector, compute GCD for u(x), vec - s for 0 <= s < p.
	// The result gives a nontrivial factorization of u.
	// The vector that results in r factors is the one
	solutions := make([][]int64, numSolutions)

	for _, vec := range vecs[1:] {
		foundSolutions := 0
		for s := int64(0); s < char; s++ {
			unit := []int64{char - s}
			p1 := PolynomialAdd(vec, unit, char)
			p := PolynomialGcd(p1, f, char)
			if len(p) != 1 {
				solutions[foundSolutions] = p
				foundSolutions++
			}
		}

		if foundSolutions == numSolutions {
			return solutions
		}
	}

	// irreducible polynomial
	return [][]int64{}
}
