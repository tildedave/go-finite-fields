package finitefields

import (
	"fmt"
	"math/rand"
	"testing"

	"github.com/stretchr/testify/assert"
)

var _ = fmt.Println

func multiplySolutions(solutions [][]int64, char int64) []int64 {
	v := []int64{1}
	for _, solution := range solutions {
		v = PolynomialMultiply(v, solution, 3)
	}

	return v
}
func TestFactorBerlekamp(t *testing.T) {
	u := []int64{8, 2, 8, 10, 10, 0, 1, 0, 1}

	solutions := FactorBerlekamp(u, 13)
	assert.Equal(t, 3, len(solutions))
	assert.Contains(t, solutions, []int64{3, 1})
	assert.Contains(t, solutions, []int64{12, 4, 8, 1})
	assert.Contains(t, solutions, []int64{6, 4, 3, 2, 1})
}

func TestFactorBerlekamp2(t *testing.T) {
	u := make([]int64, 17)
	u[16] = 1
	u[1] = 1

	// These are all the irreducibles of degree <= 4
	solutions := FactorBerlekamp(u, 2)
	assert.Equal(t, 6, len(solutions))
	assert.Contains(t, solutions, []int64{1, 1, 1, 1, 1})
	assert.Contains(t, solutions, []int64{1, 1, 0, 0, 1})
	assert.Contains(t, solutions, []int64{1, 0, 0, 1, 1})
	assert.Contains(t, solutions, []int64{0, 1})
	assert.Contains(t, solutions, []int64{1, 1})
	assert.Contains(t, solutions, []int64{1, 1, 1})
	// assert.Contains(t, solutions, []int64{6, 4, 3, 2, 1})
}

func countByDegree(solutions [][]int64, degree int) int {
	numResults := 0
	for _, solution := range solutions {
		if PolynomialDegree(solution) == degree {
			numResults++
		}
	}

	return numResults
}

func TestFactorBerlekamp3(t *testing.T) {
	u := make([]int64, 65)
	u[64] = 1
	u[1] = 1

	// These are all the irreducibles of degree <= 6
	solutions := FactorBerlekamp(u, 2)
	// degree 1 (2 irreducibles) degree 2 (1), degree 3 (2) degree 6 (9)
	assert.Equal(t, 9+2+1+2, len(solutions))

	assert.Equal(t, 2, countByDegree(solutions, 1))
	assert.Equal(t, 1, countByDegree(solutions, 2))
	assert.Equal(t, 2, countByDegree(solutions, 3))
	assert.Equal(t, 9, countByDegree(solutions, 6))
}

func TestFactorBerlekamp4(t *testing.T) {
	u := make([]int64, 28)
	u[27] = 1
	u[1] = 2

	solutions := FactorBerlekamp(u, 3)
	assert.Equal(t, 11, len(solutions))
	assert.Contains(t, solutions, []int64{0, 1})
	assert.Contains(t, solutions, []int64{1, 1})
	assert.Contains(t, solutions, []int64{2, 1})
	assert.Contains(t, solutions, []int64{1, 2, 0, 1})
	assert.Contains(t, solutions, []int64{2, 2, 0, 1})
	assert.Contains(t, solutions, []int64{2, 0, 1, 1})
	assert.Contains(t, solutions, []int64{2, 1, 1, 1})
	assert.Contains(t, solutions, []int64{1, 2, 1, 1})
	assert.Contains(t, solutions, []int64{1, 0, 2, 1})
	assert.Contains(t, solutions, []int64{1, 1, 2, 1})
	assert.Contains(t, solutions, []int64{2, 2, 2, 1})

	assert.Equal(t, u, multiplySolutions(solutions, 3))
}

func TestFactorBerlekampIrreducible(t *testing.T) {
	p := []int64{1, 1, 0, 0, 1}
	solutions := FactorBerlekamp(p, 2)
	assert.Equal(t, 0, len(solutions))
}

func TestFactorDistinctDegree(t *testing.T) {
	u := []int64{8, 2, 8, 10, 10, 0, 1, 0, 1}

	solutions := FactorDistinctDegree(u, 13)
	assert.Equal(t, 3, len(solutions))

	polynomials := make([][]int64, len(solutions))
	i := 0

	for _, solution := range solutions {
		polynomials[i] = solution.factor
		i++
	}
	assert.Contains(t, polynomials, []int64{3, 1})
	assert.Contains(t, polynomials, []int64{12, 4, 8, 1})
	assert.Contains(t, polynomials, []int64{6, 4, 3, 2, 1})
}

func TestFactorDistinctDegree2(t *testing.T) {
	u := make([]int64, 65)
	u[64] = 1
	u[1] = 1

	solutions := FactorDistinctDegree(u, 2)

	v := []int64{1}
	for _, solution := range solutions {
		v = PolynomialMultiply(v, solution.factor, 2)
	}
	assert.Equal(t, v, u)
}

func TestFactorCantorZassenhaus(t *testing.T) {
	r := rand.New(rand.NewSource(0))

	u := make([]int64, 28)
	u[27] = 1
	u[1] = 2

	solutions := FactorCantorZassenhaus(r, u, 3)

	assert.Equal(t, 11, len(solutions))
	assert.Contains(t, solutions, []int64{0, 1})
	assert.Contains(t, solutions, []int64{1, 1})
	assert.Contains(t, solutions, []int64{2, 1})
	assert.Contains(t, solutions, []int64{1, 2, 0, 1})
	assert.Contains(t, solutions, []int64{2, 2, 0, 1})
	assert.Contains(t, solutions, []int64{2, 0, 1, 1})
	assert.Contains(t, solutions, []int64{2, 1, 1, 1})
	assert.Contains(t, solutions, []int64{1, 2, 1, 1})
	assert.Contains(t, solutions, []int64{1, 0, 2, 1})
	assert.Contains(t, solutions, []int64{1, 1, 2, 1})
	assert.Contains(t, solutions, []int64{2, 2, 2, 1})

	assert.Equal(t, u, multiplySolutions(solutions, 3))
}
