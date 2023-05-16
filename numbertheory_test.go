package finitefields

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
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
	assert.False(t, IsPrime(16))
	assert.False(t, IsPrime(25))
}

func TestPrimeDivisors(t *testing.T) {
	assert.Equal(t, []int64{2, 5, 5}, PrimeDivisors(50))
	assert.Equal(t, []int64{2, 3}, PrimeDivisors(6))
}

func TestMobiusInversion(t *testing.T) {
	assert.Equal(t, 1, MobiusInversion(1))
	assert.Equal(t, -1, MobiusInversion(2))
	assert.Equal(t, -1, MobiusInversion(3))
	assert.Equal(t, 0, MobiusInversion(4))
	assert.Equal(t, -1, MobiusInversion(5))
	assert.Equal(t, 1, MobiusInversion(6))
	assert.Equal(t, -1, MobiusInversion(7))
	assert.Equal(t, 0, MobiusInversion(8))
	assert.Equal(t, 0, MobiusInversion(9))
}
