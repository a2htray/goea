// Package math provides the operations among vector or matrix with
// underlying []float64 structure
package math

import "errors"

var (
	errorLengthNotEqual = errors.New("the lengths are not equal")
	errorDividerZero = errors.New("divider zero")
)

// Vector
type Vector []float64

// N returns the length of a vector
func (v Vector) N() int {
	return len(v)
}

// Copy uses a deep copy from an origin vector
func (v Vector) Copy() Vector {
	ret := make(Vector, v.N())
	copy(ret, v)
	return ret
}

// Float64 ...
func (v Vector) Float64() []float64 {
	return []float64(v)
}

// Add method would change the origin vector
// the origin vector will be changed to the sum vector of two given vectors
func (v Vector) Add(vector Vector) error {
	if v.N() != vector.N() {
		return errorLengthNotEqual
	}

	for i := 0; i < v.N(); i++ {
		v[i] += vector[i]
	}
	return nil
}

// Sub method would change the origin vector
// the origin vector will be subtracted by the given vector
func (v Vector) Sub(vector Vector) error {
	if v.N() != vector.N() {
		return errorLengthNotEqual
	}

	for i := 0; i < v.N(); i++ {
		v[i] -= vector[i]
	}
	return nil
}

// Multiply method will change the origin vector
func (v Vector) Multiply(factor float64) {
	for i := 0; i < v.N(); i++ {
		v[i] *= factor
	}
}

// Divide method will change the origin vector
func (v Vector) Divide(divider float64) error {
	if divider == 0 {
		return errorDividerZero
	}

	for i := 0; i < v.N(); i++ {
		v[i] /= divider
	}
	return nil
}

// VectorAdd ...
func VectorAdd(v1, v2 Vector) (Vector, error) {
	if v1.N() != v2.N() {
		return nil, errorLengthNotEqual
	}
	ret := v1.Copy()
	for i := 0; i < v1.N(); i++ {
		ret[i] += v1[i]
	}
	return ret, nil
}

// VectorSub ...
func VectorSub(v1, v2 Vector) (Vector, error) {
	if v1.N() != v2.N() {
		return nil, errorLengthNotEqual
	}
	ret := v1.Copy()
	for i := 0; i < v1.N(); i++ {
		ret[i] -= v2[i]
	}
	return ret, nil
}

// VectorMultiply ...
func VectorMultiply(v Vector, factor float64) Vector {
	v1 := v.Copy()
	for i := 0; i < v.N(); i++ {
		v1[i] *= factor
	}
	return v1
}

// VectorDivide ...
func VectorDivide(v Vector, divider float64) (Vector, error) {
	if divider == 0 {
		return nil, errorDividerZero
	}
	v1 := v.Copy()
	for i := 0; i < v.N(); i++ {
		v1[i] /= divider
	}
	return v1, nil
}


