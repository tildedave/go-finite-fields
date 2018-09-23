package main

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

var _ = fmt.Println

func TestPolynomialToString(t *testing.T) {
	assert.Equal(t, PolynomialToString([]int64{1, 2, 3}), "3x^2 + 2x + 1")
	assert.Equal(t, PolynomialToString([]int64{4, 5, 1}), "x^2 + 5x + 4")
}

func TestPolynomialAdd(t *testing.T) {
	assert.Equal(t, []int64{5, 7, 9}, PolynomialAdd([]int64{1, 2, 3}, []int64{4, 5, 6}, 13))
	assert.Equal(t, []int64{1, 2, 6, 4}, PolynomialAdd([]int64{1, 0, 0, 4}, []int64{0, 2, 6}, 13))
}

func TestPolynomialMult(t *testing.T) {
	assert.Equal(t, []int64{3, 3, 3}, PolynomialMultiply([]int64{3}, []int64{1, 1, 1}, 5))
	assert.Equal(t, []int64{1, 4, 6, 4, 1}, PolynomialMultiply([]int64{1, 2, 1}, []int64{1, 2, 1}, 5))
}

func TestModInverse(t *testing.T) {
	for i := int64(1); i < 7; i++ {
		assert.Equal(t, ((ModInverse(i, 7) * i) % 7), int64(1))
	}
}

func TestPolynomialMod(t *testing.T) {
	assert.Equal(t, []int64{1, 2, 0}, PolynomialMod([]int64{4, 14, 6}, 3))
}

func TestPolynomialTrunc(t *testing.T) {
	assert.Equal(t, []int64{1, 2}, PolynomialTrunc([]int64{1, 2, 0, 0, 0, 0}))
}
