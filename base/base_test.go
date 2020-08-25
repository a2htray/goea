package base

import (
	"fmt"
	"github.com/a2htray/goea"
	"testing"
)

func TestNewEAModel(t *testing.T) {
	m := 5
	n := 2
	iterNum := 10
	upper := Boundary{10, 10}
	lower := Boundary{-10, -10}

	model := NewEAModel(m, n, iterNum,true, Limit{
		Upper: upper,
		Lower: lower,
	}, goea.ObjectGriewangk)
	fmt.Println("once")
	for i := 0; i < m; i++ {
		fmt.Println(model.Population[i].Vector, model.Population[i].FitnessValue)
	}
	fmt.Println("twice")
	model.ComputeCurrentFitnessValues()
	for i := 0; i < m; i++ {
		fmt.Println(model.Population[i].Vector, model.Population[i].FitnessValue)
	}

	fmt.Println("best individual:")
	fmt.Println(model.GlobalBestIndividual)

}