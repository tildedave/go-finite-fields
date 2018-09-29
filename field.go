package main

import (
	"fmt"
)

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

func Max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}

func ModExp(m int64, n int64, p int64) int64 {
	var pow int64 = 1

	for n > 0 {
		if n%2 == 1 {
			pow = (pow * m) % p
		}
		m = (m * m) % p
		n = n / 2
	}

	return pow
}

func ModInverse(x int64, p int64) int64 {
	// a^{p-1} = 1 mod p
	// a * a^{p-2} = 1 mod p
	return ModExp(x, p-2, p)
}

func PolynomialDegree(f1 []int64) int {
	return len(f1) - 1
}

func PolynomialAdd(f1 []int64, f2 []int64, char int64) []int64 {
	// assume deg(f1) <= f2
	l1 := len(f1)
	l2 := len(f2)
	if l1 > l2 {
		temp := f1
		f1 = f2
		f2 = temp
		l1 = len(f1)
		l2 = len(f2)
	}

	add := make([]int64, l2)
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

func PolynomialDivide(f []int64, g []int64, char int64) ([]int64, []int64) {
	// return f1 / f2
	// for now assume deg(f1 <= f2)

	// d1 := PolynomialDegree(f1)
	// synthetic division
	out := make([]int64, len(g))
	copy(out, g)

	inv := ModInverse(f[len(f)-1], char)

	for i := len(g) - 1; i >= len(f)-1 && i >= 0; i-- {
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

func PolynomialMod(f []int64, g []int64, char int64) []int64 {
	_, r := PolynomialDivide(f, g, char)
	return r
}

type Field struct {
	characteristic int64
}
