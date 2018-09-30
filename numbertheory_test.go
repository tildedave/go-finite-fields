package main

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

var _ = fmt.Println

func TestModInverse(t *testing.T) {
	for i := int64(1); i < 7; i++ {
		assert.Equal(t, ((ModInverse(i, 7) * i) % 7), int64(1))
	}
}

func TestISqrt(t *testing.T) {
	assert.Equal(t, int64(1), isqrt(1))
	assert.Equal(t, int64(3), isqrt(10))
	assert.Equal(t, int64(4), isqrt(20))
	assert.Equal(t, int64(5), isqrt(25))
}

func TestIsPrime(t *testing.T) {
	assert.False(t, IsPrime(1))
	assert.True(t, IsPrime(2))
	assert.True(t, IsPrime(3))
	assert.False(t, IsPrime(4))
	assert.True(t, IsPrime(5))
	assert.False(t, IsPrime(6))
	assert.True(t, IsPrime(7))
	assert.False(t, IsPrime(8))
	assert.False(t, IsPrime(9))
	assert.False(t, IsPrime(10))
	assert.True(t, IsPrime(11))
	assert.False(t, IsPrime(12))
	assert.True(t, IsPrime(13))
}
