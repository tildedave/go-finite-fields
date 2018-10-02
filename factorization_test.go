package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

var _ = fmt.Println

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

	// These are all the irreducibles of degree 4
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

func TestFactorBerlekampIrreducible(t *testing.T) {
	p := []int64{1, 1, 0, 0, 1}
	solutions := FactorBerlekamp(p, 2)
	assert.Equal(t, 0, len(solutions))
}
