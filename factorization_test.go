package main

import (
	"fmt"
	"testing"
)

var _ = fmt.Println

func TestFactorBerlekamp(t *testing.T) {
	u := []int64{8, 2, 8, 10, 10, 0, 1, 0, 1}

	FactorBerlekamp(u, 13)
}
