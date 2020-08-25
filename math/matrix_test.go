package math

import (
	"testing"
)

func TestMatrix_Size(t *testing.T) {
	mat := Matrix{
		Vector{1, 2, 3, 4},
	}

	m, n := mat.Size()

	if m != 1 || n != 4 {
		t.Fatal("m != 1 || n != 4, error")
	}
}

func TestMatrix_Get(t *testing.T) {
	v1 := Vector{1, 2, 3}
	mat := Matrix{
		v1,
	}
	defer func() {
		if err := recover(); err != nil && err != errorIndexOutOfRange {
			t.Log(err)
			t.Fatal("err != nil && err != errorIndexOutOfRange, err")
		}
	}()
	copyV1, err := mat.Get(0, true)
	if err != nil {
		panic(err)
	}
	if &v1 == &copyV1 {
		t.Fatal("&v1 == &copyV1, error")
	}

	originV1, err := mat.Get(0, false)
	if err != nil {
		panic(err)
	}
	if &v1 == &originV1 {
		t.Fatal("&v1 == &originV1, error")
	}
}