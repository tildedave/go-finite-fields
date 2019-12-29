package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestExp(t *testing.T) {
	assert.Equal(t, int64(32), IntExp(2, 5))
	assert.Equal(t, int64(512), IntExp(2, 9))
	assert.Equal(t, int64(32768), IntExp(2, 15))
}
