package main

// From https://github.com/high-code-coverage/golang-example/blob/master/math/gcd.go#L3
// (For now)
func GCD(a, b int64) int64 {
	for a != b {
		if a > b {
			a -= b
		} else {
			b -= a
		}
	}

	return a
}
