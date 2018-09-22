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
	assert.Equal(t, []int64{5, 7, 9}, PolynomialAdd([]int64{1, 2, 3}, []int64{4, 5, 6}))
	assert.Equal(t, []int64{1, 2, 6, 4}, PolynomialAdd([]int64{1, 0, 0, 4}, []int64{0, 2, 6}))
}

func TestPolynomialMod(t *testing.T) {
	assert.Equal(t, []int64{1, 2, 0}, PolynomialMod([]int64{4, 14, 6}, 3))
}
