package goea

import (
	"fmt"
	"math/rand"
	"testing"
	"time"
)

func TestApplyTo(t *testing.T) {
	fc := func(vector []float64) float64 {
		return vector[0] + vector[1]
	}

	individual := Individual{1, 2}

	if individual.ApplyTo(fc) != 3 {
		t.Fatal("the result must be 3")
	}
}

func TestApplyTo2(t *testing.T) {
	population := Population{
		Individual{1, 2},
		Individual{3, 4},
		Individual{5, 6},
	}

	fitnessCollection := population.ApplyTo(func(vector []float64) float64 {
		return vector[0] + vector[1]
	})

	if fitnessCollection[0] != 3 || fitnessCollection[1] != 7 || fitnessCollection[2] != 11 {
		t.Fatal(`fitnessCollection[0]=3,fitnessCollection[1]=7,fitnessCollection[2]=11`)
	}
}

func TestInitPopulation(t *testing.T) {
	population := initPopulation(5, 2, Boundary{
		Lower: []float64{1, 1, 1, 1, 1, 1},
		Upper: []float64{2, 2, 2, 2, 2, 2},
	})

	t.Log(population)
}

func TestInitPopulation2(t *testing.T) {
	defer func() {
		if r := recover(); r != errorNumOfIndividuals && r != errorNumOfComponents {
			t.Fatal("this is a wrong test")
		}
	}()
	population := initPopulation(-1, -1, Boundary{
		Lower: []float64{1, 1, 1, 1, 1, 1},
		Upper: []float64{2, 2, 2, 2, 2, 2},
	})

	t.Log(population)
}
func TestNewIndividual(t *testing.T) {
	boundary := Boundary{
		Lower: []float64{1, 1, 1, 1, 1, 1, 1, 1, 1},
		Upper: []float64{2, 2, 2, 2, 2, 2, 2, 2, 2},
	}

	t.Log(NewIndividual(9, boundary))
}

func TestRand(t *testing.T) {
	for i := 0; i < 10; i++ {
		rand.Seed(time.Now().UnixNano())
		t.Log(rand.Int())
	}

	rng := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < 10; i++ {
		t.Log(rng.Int())
	}
}

func TestIndividual_String(t *testing.T) {
	individual := NewIndividual(4, Boundary{
		Lower: []float64{1, 1, 1, 1, 1, 1, 1, 1, 1},
		Upper: []float64{2, 2, 2, 2, 2, 2, 2, 2, 2},
	})
	t.Log(individual.String())
}

func TestNewIndividual2(t *testing.T) {
	defer func() {
		if r := recover(); r != errorBoundaryNumNotThanTo {
			t.Fatal("this is a wrong test")
		}
	}()
	boundary := Boundary{
		Lower: []float64{1, 1, 1, 1, 1, 1, 1, 1, 1},
		Upper: []float64{2, 2, 2, 2, 2, 2, 2, 2, 2},
	}

	NewIndividual(100, boundary)
}

func TestPopulation_String(t *testing.T) {
	population := initPopulation(5, 2, Boundary{
		Lower: []float64{1, 1, 1, 1, 1, 1},
		Upper: []float64{2, 2, 2, 2, 2, 2},
	})

	t.Log(population.String())
}

func TestVector_Mat(t *testing.T) {
	v := Vector{
		1, 2, 3,
	}
	t.Log(v.Mat())
}

func TestDivide(t *testing.T) {
	fmt.Println(9 / 3)
	fmt.Println(10 / 3)
}

func TestIndividual_Copy(t *testing.T) {
	src := Individual{1, 2, 3}
	dest := src.Copy()

	for i, v := range src {
		if v != dest[i] {
			t.Fatal("not equal")
		}
	}

	if &src == &dest {
		t.Fatal("the pointer should not be equal")
	}
}
