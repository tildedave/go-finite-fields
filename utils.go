package main

// From https://github.com/high-code-coverage/golang-example/blob/master/math/gcd.go#L3
// (For now)
func GCD(a, b int64) int64 {
	if a < 0 {
		return -GCD(-a, b)
	}
	if b < 0 {
		return -GCD(a, -b)
	}
	for a != b {
		if a > b {
			a -= b
		} else {
			b -= a
		}
	}

	return a
}

func IntExp(a int64, n int64) int64 {
	pow := int64(1)

	for n > 0 {
		if n%2 == 1 {
			pow = pow * a
		}
		a = a * a
		n = n / 2
	}

	return pow
}
