package finitefields

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
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

	assert.Equal(t, []int64{6, 3, 1, 3}, PolynomialMultiply(p1, p2, 7))
}

func TestPolynomialScalarMod(t *testing.T) {
	assert.Equal(t, []int64{1, 2, 0}, PolynomialScalarMod([]int64{4, 14, 6}, 3))
}

func TestPolynomialTrunc(t *testing.T) {
	assert.Equal(t, []int64{1, 2}, PolynomialTrunc([]int64{1, 2, 0, 0, 0, 0}))
	assert.Equal(t, []int64{2}, PolynomialTrunc([]int64{2, 0}))
}

func TestPolynomialDivide1(t *testing.T) {
	q, r := PolynomialDivide([]int64{1, 2, 1}, []int64{1, 4, 6, 4, 1}, 5)
	assert.Equal(t, []int64{1, 2, 1}, q)
	assert.Equal(t, []int64{}, r)
}

func TestPolynomialDivide2(t *testing.T) {
	q, r := PolynomialDivide([]int64{1, 2}, []int64{6, 2, 6, 3}, 7)
	assert.Equal(t, []int64{6, 4, 5}, q)
	assert.Equal(t, []int64{}, r)
}

func TestPolynomialDivide3(t *testing.T) {
	q, r := PolynomialDivide([]int64{2, 0, 1}, []int64{0, 0, 2}, 3)
	assert.Equal(t, []int64{2}, q)
	assert.Equal(t, []int64{2}, r)
}

func TestPolynomialDivide4(t *testing.T) {
	q, r := PolynomialDivide([]int64{2, 0, 1}, []int64{0, 0, 2, 1}, 3)

	assert.Equal(t, []int64{2, 1}, r)
	assert.Equal(t, []int64{2, 1}, q)
}

func TestPolynomialDerivative(t *testing.T) {
	// x^2 + 6x - 1 = 2x + 6
	assert.Equal(t, []int64{6, 2}, PolynomialDerivative([]int64{3, 6, 1}, 7))
	assert.Equal(t, []int64{3, 0, 0, 4}, PolynomialDerivative([]int64{0, 3, 0, 0, 1}, 7))
}

func TestPolynomialMakeMonic(t *testing.T) {
	assert.Equal(t, []int64{1, 1}, PolynomialMakeMonic([]int64{3, 3}, 5))
}

func TestPolynomialGcd(t *testing.T) {
	p1 := []int64{6, 2}
	p2 := []int64{3, 4, 1}
	p3 := []int64{3, 1}

	p4 := PolynomialMultiply(p1, p2, 7)
	p5 := PolynomialMultiply(p1, p3, 7)
	gcd := PolynomialGcd(p4, p5, 7)

	assert.True(t, PolynomialDivides(gcd, p4, 7))
	assert.True(t, PolynomialDivides(gcd, p5, 7))
}

func TestPolynomialGcd2(t *testing.T) {
	gcd := PolynomialGcd([]int64{193 - 6, 193 - 5, 1}, []int64{6, 7, 1}, 193)
	assert.Equal(t, []int64{1, 1}, gcd)
}

func TestPolynomialIsSquareFree(t *testing.T) {
	p := []int64{2, 1}
	p2 := PolynomialMultiply(p, p, 3)

	assert.True(t, PolynomialIsSquareFree(p, 3))
	assert.False(t, PolynomialIsSquareFree(p2, 3))
}

func TestPolynomialMod(t *testing.T) {
	// https://www.doc.ic.ac.uk/~mrh/330tutor/ch04s02.html
	modulus := []int64{2, 0, 1}
	assert.Equal(t, []int64{2}, PolynomialMod([]int64{0, 0, 2}, modulus, 3))
	assert.Equal(t, []int64{1}, PolynomialMod([]int64{0, 0, 0, 0, 1}, modulus, 3))
	assert.Equal(t, []int64{0, 1}, PolynomialMod([]int64{0, 0, 0, 1}, modulus, 3))
	assert.Equal(t, []int64{}, PolynomialMod([]int64{0, 0, 2, 0, 1}, modulus, 3))
	assert.Equal(t, []int64{2, 1}, PolynomialMod([]int64{0, 0, 2, 1}, modulus, 3))
	assert.Equal(t, []int64{2}, PolynomialMod([]int64{0, 0, 0, 0, 0, 0, 2}, modulus, 3))
}

func TestPolynomialModExp(t *testing.T) {
	// Example is from Knuth 4.6.2
	mod := []int64{8, 2, 8, 10, 10, 0, 1, 0, 1}
	x := []int64{0, 1}

	assert.Equal(t, []int64{2, 1, 7, 11, 10, 12, 5, 11}, PolynomialModExp(x, 13, mod, 13))
	assert.Equal(t, []int64{3, 6, 4, 3, 0, 4, 7, 2}, PolynomialModExp(x, 13*2, mod, 13))
	assert.Equal(t, []int64{4, 3, 6, 5, 1, 6, 2, 3}, PolynomialModExp(x, 13*3, mod, 13))
	assert.Equal(t, []int64{2, 11, 8, 8, 3, 1, 3, 11}, PolynomialModExp(x, 13*4, mod, 13))
	assert.Equal(t, []int64{6, 11, 8, 6, 2, 7, 10, 9}, PolynomialModExp(x, 13*5, mod, 13))
	assert.Equal(t, []int64{5, 11, 7, 10, 0, 11, 7, 12}, PolynomialModExp(x, 13*6, mod, 13))
	assert.Equal(t, []int64{3, 3, 12, 5, 0, 11, 9, 12}, PolynomialModExp(x, 13*7, mod, 13))
}
