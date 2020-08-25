package math

import "errors"

type Matrix []Vector

var (
	errorIndexOutOfRange = errors.New("index out of range")
)

// Size return the size of a matrix
func (m Matrix) Size() (m1, n int) {
	m1 = len(m)
	if m1 == 0 {
		n = 0
	} else {
		n = m[0].N()
	}
	return
}

// Get the ith vector in a matrix
func (m Matrix) Get(i int, copy bool) (Vector, error) {
	m1, _ := m.Size()
	if i < 0 || i >= m1 {
		return nil, errorIndexOutOfRange
	}

	if copy {
		return m[i].Copy(), nil
	} else {
		return m[i], nil
	}
}