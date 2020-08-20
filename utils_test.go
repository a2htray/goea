package goea

import (
	"fmt"
	"math/rand"
	"testing"
	"time"
)

func TestSubtract(t *testing.T) {
	a := []float64{1, 2, 3, 4}
	b := []float64{1, 2, 3, 4}

	result := Subtract(a, b)

	if result[0] != 0 || result[1] != 0 || result[2] != 0 || result[3] != 0 {
		t.Fatal("result[0]=0, result[1]=0, result[2]=0, result[3]=0")
	}
}

func TestExpandX(t *testing.T) {
	mat1 := Mat{
		[]float64{1, 2},
		[]float64{3, 4},
	}
	mat2 := Mat{
		[]float64{9},
		[]float64{9},
	}
	t.Log(MatExpandX(mat1, mat2))
}

func TestVectorExpand(t *testing.T) {
	v1 := Vector{1, 2, 3}
	v2 := Vector{3, 2, 1}
	t.Log(VectorExpand(v1, v2))
}

func TestRandIntRange(t *testing.T) {
	rng := rand.New(rand.NewSource(time.Now().Unix()))
	for i := 0; i < 10; i++ {
		r := RandIntRange(1, 100, rng)
		fmt.Println(r)
	}
}

func TestRangeInt(t *testing.T) {
	fmt.Println(RangeInt(0, 10))
}

func TestShuffleSliceInt(t *testing.T) {
	rng := rand.New(rand.NewSource(time.Now().Unix()))
	fmt.Println(ShuffleSliceInt(RangeInt(0, 10), rng))
}
