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
	assert.Equal(t, []int64{1, 4, 6, 4, 1}, PolynomialMultiply([]int64{1, 2, 1}, []int64{1, 2, 1}, 7))

	p1 := []int64{1, 2}
	p2 := []int64{6, 5, 5}
	fmt.Printf("multiply %s by %s\n", PolynomialToString(p1), PolynomialToString(p2))
	fmt.Println(PolynomialToString(PolynomialMultiply(p1, p2, 7)))
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
	assert.Equal(t, []int64{2}, PolynomialTrunc([]int64{2, 0}))
}

func TestPolynomialDivide(t *testing.T) {
	q, r := PolynomialDivide([]int64{1, 2, 1}, []int64{1, 4, 6, 4, 1}, 5)
	assert.Equal(t, []int64{1, 2, 1}, q)
	assert.Equal(t, []int64{}, r)

	q, r = PolynomialDivide([]int64{1, 2}, []int64{6, 2, 6, 3}, 7)
	assert.Equal(t, []int64{6, 4, 5}, q)
	assert.Equal(t, []int64{}, r)

	q, r = PolynomialDivide([]int64{2, 0, 1}, []int64{0, 0, 2}, 3)
	assert.Equal(t, []int64{2}, q)
	assert.Equal(t, []int64{2}, r)
}
