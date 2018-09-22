package main

import (
	"fmt"
)

func PolynomialToString(coeffs []int64) string {
	str := ""
	for i := len(coeffs) - 1; i >= 0; i-- {
		n := coeffs[i]

		if i != len(coeffs)-1 {
			str += " + "
		}

		coeff := fmt.Sprintf("%d", n)
		if n == 0 {
			continue
		} else if n == 1 {
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

	for n > 1 {
		if n%2 == 0 {
			pow = (pow * m * m) % p
			n = n / 2
		} else {
			pow = (pow * m) % p
			n = n - 1
		}
	}

	return pow
}

func ModInverse(x int64, p int64) int64 {
	// a^{p-1} = 1 mod p
	// a * a^{p-2} = 1 mod p
	return ModExp(x, p-2, p)
}

func PolynomialDegree(f1 []int64) int {
	return len(f1)
}

func PolynomialAdd(f1 []int64, f2 []int64, char int64) []int64 {
	// assume deg(f1) <= f2
	d1 := PolynomialDegree(f1)
	d2 := PolynomialDegree(f2)
	if d1 > d2 {
		temp := f1
		f1 = f2
		f2 = temp
		d1 = PolynomialDegree(f1)
		d2 = PolynomialDegree(f2)
	}

	add := make([]int64, d2)
	copy(add, f2)
	for i := 0; i < d1; i++ {
		add[i] += f1[i] % char
	}

	return add
}

func PolynomialMod(f []int64, char int64) []int64 {
	g := make([]int64, len(f))
	copy(g, f)
	for i, x := range g {
		g[i] = x % char
	}

	return g
}

func PolynomialDivide(f1 []int64, f2 []int64, char int64) {
	// return f1 / f2

	// for now assume deg(f1 <= f2)

}

type Field struct {
	characteristic int64
}
