package finitefields

import (
	"errors"
	"fmt"
)

type Polynomial struct {
	coeffs []int64
	field  *Field
}

type Field struct {
	characteristic int64
	polynomial     *Polynomial
}

func MakePolynomial(coeffs []int64, field *Field) Polynomial {
	return Polynomial{coeffs: coeffs, field: field}
}

func FiniteField(characteristic int64) (*Field, error) {
	if !IsPrime(characteristic) {
		msg := fmt.Sprintf("Cannot create a field with non-prime characteristic: %d", characteristic)
		return nil, errors.New(msg)
	}
	f := Field{characteristic: characteristic, polynomial: nil}
	return &f, nil
}
