package main

func Max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}

func isqrt(p int64) int64 {
	if p < 2 {
		return p
	}

	small := isqrt(p>>2) << 1
	large := small + 1
	if large*large > p {
		return small
	}

	return large
}

func IsPrime(p int64) bool {
	if p == 2 || p == 3 {
		return true
	}

	mod6 := p % 6
	if mod6 != 1 && mod6 != 5 {
		return false
	}

	if p == 1 {
		return false
	}

	for i := int64(3); i <= isqrt(p); i++ {
		if p%i == 0 {
			return false
		}
	}

	return true
}

// PrimeDivisors returns the prime divisors of the number with multiplicity.
func PrimeDivisors(n int64) []int64 {
	divisors := make([]int64, n)
	numDivisors := 0

	for i := int64(2); i <= n; i++ {
		if n%i != 0 {
			// not a divisor
			continue
		}

		for n%i == 0 {
			n = n / i
			divisors[numDivisors] = i
			numDivisors++
		}
	}

	if numDivisors == 0 {
		divisors[numDivisors] = n
		numDivisors++
	}

	return divisors[0:numDivisors]
}

// MobiusInversion returns:
// 1 if the number = 1
// 0 if the number contains a square divisor
// 1 if the number contains an even number of prime divisors
// -1 if the number contains an odd number of prime divisors
func MobiusInversion(n int64) int {
	if n == 1 {
		return 1
	}

	divisors := PrimeDivisors(n)
	for i := 0; i < len(divisors)-1; i++ {
		if divisors[i] == divisors[i+1] {
			// number is a square
			return 0
		}
	}

	if len(divisors)%2 == 0 {
		return 1
	}

	return -1
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
