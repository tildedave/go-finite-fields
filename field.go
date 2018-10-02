package main

import (
	"errors"
	"fmt"
)

type Polynomial struct {
	coeffs []int64
	field  *Field
}

type Field struct {
	characteristic int64
	polynomial     *Polynomial
}

func MakePolynomial(coeffs []int64, field *Field) Polynomial {
	return Polynomial{coeffs: coeffs, field: field}
}

func FiniteField(characteristic int64) (*Field, error) {
	if !IsPrime(characteristic) {
		msg := fmt.Sprintf("Cannot create a field with non-prime characteristic: %d", characteristic)
		return nil, errors.New(msg)
	}
	f := Field{characteristic: characteristic, polynomial: nil}
	return &f, nil
}

func PolynomialToString(coeffs []int64) string {
	str := ""

	for i := len(coeffs) - 1; i >= 0; i-- {
		n := coeffs[i]

		if n == 0 {
			continue
		}

		if i != len(coeffs)-1 {
			str += " + "
		}

		coeff := fmt.Sprintf("%d", n)
		if n == 1 {
			coeff = ""
		}

		if i == 0 {
			if n == 1 {
				str += "1"
			} else {
				str += coeff
			}
		} else if i == 1 {
			str += fmt.Sprintf("%sx", coeff)
		} else {
			str += fmt.Sprintf("%sx^%d", coeff, i)
		}
	}

	return str
}

func PolynomialDegree(f1 []int64) int {
	return len(f1) - 1
}

func PolynomialAdd(f1 []int64, f2 []int64, char int64) []int64 {
	if len(f1) > len(f2) {
		return PolynomialAdd(f2, f1, char)
	}

	add := make([]int64, len(f2))
	copy(add, f2)
	for i, x := range f1 {
		add[i] = (add[i] + x) % char
	}

	return add
}

func PolynomialScalarMod(f []int64, char int64) []int64 {
	g := make([]int64, len(f))
	copy(g, f)
	for i, x := range g {
		g[i] = x % char
	}

	return g
}

func PolynomialMultiply(f []int64, g []int64, char int64) []int64 {
	// dimension of f * g = deg(f) * deg(g).  len(f) = deg(f) + 1
	d1 := len(f) - 1
	d2 := len(g) - 1
	dim := d1 + d2
	result := make([]int64, dim+1)

	// 0 term = f[0] * g[0]
	// 1 term = f[1] * g[0] + f[0] * g[1]
	// 2 term = f[2] * g[0] + f[1] * g[1] + f[0] * g[2]

	for n := 0; n <= dim; n++ {
		for i := 0; i <= n; i++ {
			j := n - i
			if i < len(f) && j < len(g) {
				result[n] = (result[n] + f[i]*g[j]) % char
			}
		}
	}

	return result
}

func PolynomialScalar(f []int64, c int64, char int64) []int64 {
	q := make([]int64, len(f))
	copy(q, f)

	for i, x := range f {
		q[i] = (x * c) % char
	}

	return q
}

func PolynomialTrunc(f []int64) []int64 {
	i := len(f) - 1
	for i > 0 && f[i] == 0 {
		i--
	}
	if i == 0 && f[i] == 0 {
		return []int64{}
	}

	return f[0 : i+1]
}

// PolynomialDivide computes q, r such that g = f * q + r.
//
// It uses the method of "synthetic division".
func PolynomialDivide(f []int64, g []int64, char int64) ([]int64, []int64) {
	// return f1 / f2
	// for now assume deg(f1 <= f2)
	if len(g) < len(f) {
		// q = 0, r = g
		r := make([]int64, len(g))
		copy(r, g)
		return []int64{}, r
	}

	// d1 := PolynomialDegree(f1)
	// synthetic division
	out := make([]int64, len(g))
	copy(out, g)

	inv := ModInverse(f[len(f)-1], char)

	for i := len(g) - 1; i >= len(f)-1; i-- {
		out[i] = (out[i] * inv) % char
		x := out[i]

		for j := len(f) - 2; j >= 0; j-- {
			// Modify elements of g from top down
			// Total number of elements to modify is total number of divisor elements
			// x^4  + 4x^3 + 6x^2 + 4x + 1 divide by x^2 + 2x + 1
			// Given j in the divisor, we are going to modify ...
			// For i = 4 (length of x^4 term), we modify terms 3 2 in out.
			//     j = 1, j = 0
			// For i = 3 we modify terms 2 1 in out.
			// For i = 2 we modify terms 1 0 in out.
			// Then we're done.

			term := i - ((len(f) - 1) - j)
			y := (x * f[j]) % char
			out[term] = (out[term] + (char - y)) % char
		}
	}

	q, r := out[len(f)-1:], PolynomialTrunc(out[:len(f)-1])
	return q, r
}

// PolynomialMod computes the polynomial f modulo the polynomial g.
//
// Example: x^2 + x + 1 = x + 1 mod x^2.
func PolynomialMod(g []int64, f []int64, char int64) []int64 {
	_, r := PolynomialDivide(f, g, char)
	return r
}

// PolynomialModExp computes the modular exponent of f to the nth power,
// modulo a given polynomial.
//
// This will give a correct but stupid way to compute the rows of Berlekamp's algorithm.
// Potentially we will want to invoke the modulus after each step of the algorithm to
// avoid things like x^96 creating giant slices that we then waste a bunch of time
// dividing.
func PolynomialModExp(f []int64, n int64, mod []int64, char int64) []int64 {
	pow := []int64{1}

	for n > 0 {
		if n%2 == 1 {
			pow = PolynomialMultiply(pow, f, char)
			// TODO: do we need to modulus now?
		}
		f = PolynomialMultiply(f, f, char)
		n = n / 2
	}

	return PolynomialMod(pow, mod, char)
}

// PolynomialDivides returns whether or not f is a factor of g.
func PolynomialDivides(f []int64, g []int64, char int64) bool {
	return len(PolynomialMod(g, f, char)) == 0
}

// PolynomialDerivative computes the symbolic derivative of f.
//
// It is primarily used to determine whether or not a polynomial f is square-free.  If f
// is square-free, gcd(f, f') = 1.  If f has a factor v^2, then gcd(f, f') = v.
func PolynomialDerivative(f []int64, char int64) []int64 {
	out := make([]int64, len(f)-1)

	for i := int64(len(f) - 1); i >= 1; i-- {
		out[i-1] = (f[i] * i) % char
	}

	return out
}

func PolynomialMakeMonic(f []int64, char int64) []int64 {
	out := make([]int64, len(f))
	inv := ModInverse(f[len(f)-1], char)
	for i := len(f) - 1; i >= 0; i-- {
		out[i] = (f[i] * inv) % char
	}

	return out
}

func PolynomialGcd(f []int64, g []int64, char int64) []int64 {
	if len(f) > len(g) {
		return PolynomialGcd(g, f, char)
	}

	_, r := PolynomialDivide(f, g, char)

	if len(r) == 0 {
		return PolynomialMakeMonic(f, char)
	}

	return PolynomialGcd(r, f, char)
}

func PolynomialIsSquareFree(f []int64, char int64) bool {
	deriv := PolynomialDerivative(f, char)
	gcd := PolynomialGcd(f, deriv, char)

	return len(PolynomialTrunc(gcd)) == 1
}

// PolynomialsAreEqual returns true if the two polynomials are equal.
// This function does not check for leading zeroes.
func PolynomialsAreEqual(f []int64, g []int64) bool {
	if len(f) != len(g) {
		return false
	}

	for i := range f {
		if f[i] != g[i] {
			return false
		}
	}

	return true
}
