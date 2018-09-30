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

	for i := int64(3); i < isqrt(p); i++ {
		if p%i == 0 {
			return false
		}
	}

	return true
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
