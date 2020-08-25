package math

import (
	"fmt"
	"testing"
)

func TestVector_N(t *testing.T) {
	v := Vector{1, 2, 3, 4, 5}
	if v.N() != 5 {
		t.Fatal("v.N() must be 5")
	}
}

func TestVector_Copy(t *testing.T) {
	v1 := Vector{1, 2, 3, 4, 5}
	v2 := v1.Copy()

	if &v1 == &v2 {
		t.Fatal("&v1 == &v2, error")
	}
}

func TestVector_Add(t *testing.T) {
	v1 := Vector{1, 2, 3, 4, 5}
	v2 := Vector{1, 2, 3, 4}

	defer func() {
		if err := recover(); err != errorLengthNotEqual {
			t.Fatal("error must be errorLengthNotEqual")
		}
	}()

	if err := v1.Add(v2); err != nil {
		panic(err)
	}
}

func TestVector_Add2(t *testing.T) {
	v1 := Vector{1, 2, 3, 4, 5}
	v2 := Vector{1, 2, 3, 4, 5}

	defer func() {
		if err := recover(); err != nil && err != errorLengthNotEqual {

			t.Fatal("error must be errorLengthNotEqual")
		}
	}()

	if err := v1.Add(v2); err != nil {
		panic(err)
	}

	if v1[0] != 2 || v1[1] != 4 || v1[2] != 6 || v1[3] != 8 || v1[4] != 10 {
		t.Fatal("v1[0] != 2 || v1[1] != 4 || v1[2] != 6 || v1[3] != 8 || v1[4] != 10, error")
	}
}

func TestVector_Sub(t *testing.T) {
	v1 := Vector{1, 2, 3, 4, 5}
	v2 := Vector{1, 2, 3, 4}

	defer func() {
		if err := recover(); err != errorLengthNotEqual {
			t.Fatal("error must be errorLengthNotEqual")
		}
	}()

	if err := v1.Sub(v2); err != nil {
		panic(err)
	}
}

func TestVector_Sub2(t *testing.T) {
	v1 := Vector{1, 2, 3, 4, 5}
	v2 := Vector{1, 2, 3, 4, 5}

	defer func() {
		if err := recover(); err != nil && err != errorLengthNotEqual {
			t.Fatal("error must be errorLengthNotEqual")
		}
	}()

	if err := v1.Sub(v2); err != nil {
		panic(err)
	}

	if v1[0] != 0 || v1[1] != 0 || v1[2] != 0 || v1[3] != 0 || v1[4] != 0 {
		t.Fatal("v1[0] != 0 || v1[1] != 0 || v1[2] != 0 || v1[3] != 0 || v1[4] != 0, error")
	}
}

func TestVector_Multiply(t *testing.T) {
	v1 := Vector{1, 2, 3, 4, 5}
	factor := 2
	v1.Multiply(float64(factor))

	if v1[0] != 2 || v1[1] != 4 || v1[2] != 6 || v1[3] != 8 || v1[4] != 10 {
		t.Fatal("v1[0] != 2 || v1[1] != 4 || v1[2] != 6 || v1[3] != 8 || v1[4] != 10, error")
	}
}

func TestVector_Divide(t *testing.T) {
	v1 := Vector{1, 2, 3, 4, 5}
	divider := 0
	defer func() {
		if err := recover(); err != nil && err != errorDividerZero {
			t.Fatal("error must be errorDividerZero")
		}
	}()

	if err := v1.Divide(float64(divider)); err != nil {
		panic(err)
	}
}

func TestVector_Divide2(t *testing.T) {
	v1 := Vector{1, 2, 3, 4, 5}
	divider := 1
	defer func() {
		if err := recover(); err != nil && err != errorDividerZero {
			t.Fatal("error must be errorDividerZero")
		}
	}()

	if err := v1.Divide(float64(divider)); err != nil {
		panic(err)
	}

	if v1[0] != 1 || v1[1] != 2 || v1[2] != 3 || v1[3] != 4 || v1[4] != 5 {
		t.Fatal("v1[0] != 1 || v1[1] != 2 || v1[2] != 3 || v1[3] != 4 || v1[4] != 5, error")
	}
}

func TestVectorAdd(t *testing.T) {
	v1 := Vector{1, 2, 3, 4, 5}
	v2 := Vector{1, 2, 3, 4, 5}

	defer func() {
		if err := recover(); err != nil && err != errorLengthNotEqual {
			fmt.Println("the length must be equal")
		}
	}()

	sumVector, err := VectorAdd(v1, v2)

	if err != nil {
		panic(err)
	}

	if sumVector[0] != 2 || sumVector[1] != 4 || sumVector[2] != 6 || sumVector[3] != 8 || sumVector[4] != 10 {
		t.Fatal("sumVector[0] != 2 || sumVector[1] != 4 || sumVector[2] != 6 || sumVector[3] != 8 || sumVector[4] != 10, error")
	}

	if &v1 == &sumVector || &v2 == &sumVector {
		t.Fatal("&v1 == &sumVector || &v2 == &sumVector, error")
	}
}

func TestVectorSub(t *testing.T) {
	v1 := Vector{1, 2, 3, 4, 5}
	v2 := Vector{1, 2, 3, 4, 5}

	defer func() {
		if err := recover(); err != nil && err != errorLengthNotEqual {
			fmt.Println("the length must be equal")
		}
	}()

	subVector, err := VectorSub(v1, v2)

	if err != nil {
		panic(err)
	}

	if subVector[0] != 0 || subVector[1] != 0 || subVector[2] != 0 || subVector[3] != 0 || subVector[4] != 0 {
		t.Fatal("subVector[0] != 0 || subVector[1] != 0 || subVector[2] != 0 || subVector[3] != 0 || subVector[4] != 0, error")
	}

	if &v1 == &subVector || &v2 == &subVector {
		t.Fatal("&v1 == &subVector || &v2 == &subVector, error")
	}
}

func TestVectorMultiply(t *testing.T) {
	v := Vector{1, 2, 3}
	factor := float64(2)

	multiplyVector := VectorMultiply(v, factor)

	if &v == &multiplyVector {
		t.Fatal("&v == &multiplyVector, error")
	}

	if multiplyVector[0] != 2 || multiplyVector[1] != 4 || multiplyVector[2] != 6 {
		t.Fatal("multiplyVector[0] != 2 || multiplyVector[1] != 4 || multiplyVector[2] != 6, error")
	}
}

func TestVectorDivide(t *testing.T) {
	v := Vector{1, 2, 3}
	divider := float64(2)
	defer func() {
		if err := recover(); err != nil && err != errorDividerZero {
			t.Fatal("the divider is zero")
		}
	}()

	divideVector, err := VectorDivide(v, divider)

	if err != nil {
		panic(err)
	}

	if divideVector[0] != 0.5 || divideVector[1] != 1 || divideVector[2] != 1.5 {
		t.Fatal("divideVector[0] != 0.5 || divideVector[1] != 1 || divideVector[2] != 1.5, error")
	}
}