package goea

import (
	"fmt"
	"testing"
)

func TestMat_Cut(t *testing.T) {
	mat := Mat{
		Vector{1, 2, 3, 4, 5, 6},
		Vector{1, 2, 3, 4, 5, 6},
		Vector{1, 2, 3, 4, 5, 6},
		Vector{1, 2, 3, 4, 5, 6},
		Vector{1, 2, 3, 4, 5, 6},
		Vector{1, 2, 3, 4, 5, 6},
		Vector{1, 2, 3, 4, 5, 6},
	}

	newMat := mat.Cut(0, 3, 0, 3)

	if newMat[0][0] != 1 {
		t.Fatal("newMat[0][0]=1")
	}

	if newMat[2][2] != 3 {
		t.Fatal("newMat[2][2]=3")
	}

	newMat = mat.Cut(2, 4, 2, 5)

	if newMat[0][0] != 3 {
		t.Fatal("newMat[0][0]=3")
	}
	fmt.Println(newMat)
	if newMat[1][2] != 5 {
		t.Fatal("newMat[1][2]=5")
	}
}

func TestMat_Flat(t *testing.T) {
	mat := Mat{
		Vector{1, 2, 3, 4, 5, 6},
		Vector{1, 2, 3, 4, 5, 6},
	}

	if mat.Flat().N() != 12 {
		t.Fatal("mat.Flat().N()=12")
	}
}
