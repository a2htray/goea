package goea

import (
	"fmt"
	"testing"
)

func TestNewSimple(t *testing.T) {
	iterNum := 10
	de := NewDE(4, 2, Boundary{
		Lower: []float64{-10, -10, -10},
		Upper: []float64{10, 10, 10},
	}, iterNum, func(float64s []float64) float64 {
		return float64s[0] + float64s[1]
	}, DefaultDEConfig())

	de.Run()

	for i := 0; i < iterNum; i++ {
		fmt.Println(i+1, "代:", de.HistoryBestIndividuals[i], "->", de.HistoryBestFNC[i])
	}
}

func TestSimpleDE_Mutation(t *testing.T) {
	iterNum := 5
	de := NewDE(4, 3, Boundary{
		Lower: []float64{-10, -10, -10},
		Upper: []float64{10, 10, 10},
	}, iterNum, ObjectSphere, DefaultDEConfig())

	fmt.Println("outline")
	fmt.Println(de.Population.ApplyTo(ObjectSphere))
	de.Mutation()
	fmt.Println(de.Population.ApplyTo(ObjectSphere))
}

func TestSimpleDE_Run(t *testing.T) {
	SetMinimum(true)
	iterNum := 20
	de := NewDE(4, 3, Boundary{
		Lower: []float64{-10, -10, -10},
		Upper: []float64{10, 10, 10},
	}, iterNum, ObjectSphere, DefaultDEConfig())

	old := make(Population, len(de.Population))
	copy(old, de.Population)

	fmt.Println(de.Population.ApplyTo(ObjectSphere))
	de.Run()
	fmt.Println(de.Population.ApplyTo(ObjectSphere))

	for i, v := range de.Population {
		if old[i].ApplyTo(ObjectSphere) <= v.ApplyTo(ObjectSphere) {
			t.Fatal("after run, the fitness needs to be smaller")
		}
	}

	for i := 0; i < iterNum; i++ {
		fmt.Println(i+1, "代:", de.HistoryBestIndividuals[i], "->", de.HistoryBestFNC[i])
	}
}

func TestBest2bin(t *testing.T)  {
	SetMinimum(true)
	config := DefaultDEConfig()
	config.Mode = DEModeBest2bin
	iterNum := 20
	de := NewDE(5, 3, Boundary{
		Lower: []float64{-10, -10, -10},
		Upper: []float64{10, 10, 10},
	}, iterNum, ObjectSphere, config)

	old := make(Population, len(de.Population))
	copy(old, de.Population)

	fmt.Println(de.Population.ApplyTo(ObjectSphere))
	de.Run()
	fmt.Println(de.Population.ApplyTo(ObjectSphere))

	//for i, v := range de.Population {
	//	if old[i].ApplyTo(ObjectSphere) <= v.ApplyTo(ObjectSphere) {
	//		t.Fatal("after run, the fitness needs to be smaller")
	//	}
	//}

	for i := 0; i < iterNum; i++ {
		fmt.Println(i+1, "代:", de.HistoryBestIndividuals[i], "->", de.HistoryBestFNC[i])
	}
}