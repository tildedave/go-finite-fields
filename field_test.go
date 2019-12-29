package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

var _ = fmt.Println

func TestCreateField(t *testing.T) {
	f, err := FiniteField(5)

	assert.NotNil(t, f)
	assert.Nil(t, err)
}

func TestCreateFieldWithInvalidCharacteristic(t *testing.T) {
	f, err := FiniteField(4)
	assert.Nil(t, f)
	assert.Equal(t, "Cannot create a field with non-prime characteristic: 4", err.Error())
}
