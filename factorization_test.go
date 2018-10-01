package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

var _ = fmt.Println

func TestFactorBerlekamp(t *testing.T) {
	u := []int64{8, 2, 8, 10, 10, 0, 1, 0, 1}

	solutions := FactorBerlekamp(u, 13)
	assert.Equal(t, 3, len(solutions))
	assert.Contains(t, solutions, []int64{3, 1})
	assert.Contains(t, solutions, []int64{12, 4, 8, 1})
	assert.Contains(t, solutions, []int64{6, 4, 3, 2, 1})
}
