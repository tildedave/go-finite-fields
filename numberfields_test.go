package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestComputeNumberFieldSignature(t *testing.T) {
	assert.Equal(t, NumberFieldSignature{numReal: 1, numComplex: 1},
		ComputeNumberFieldSignature([]int64{2, 0, 0, 1}),
		"x^3 + 2 should have had 1 real and 1 complex embedding")
	assert.Equal(t, NumberFieldSignature{numReal: 0, numComplex: 1},
		ComputeNumberFieldSignature([]int64{1, 1, 1}),
		"x^2 + x + 1 should have had 1 complex embedding")
}
