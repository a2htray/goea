package goea

import (
	"fmt"
	"testing"
)

func TestABC_OnlookerStage(t *testing.T) {
	abc := NewABC(10, 3, Boundary{
		Lower: []float64{-10, -10, -10},
		Upper: []float64{10, 10, 10},
	}, 10, 3, func(float64s []float64) float64 {
		return float64s[0] + float64s[1] + float64s[2]
	})

	for i := 0; i < 10; i++ {
		abc.EmployeeStage()
		abc.OnlookerStage()
		abc.ScouterStage()
		t.Log(abc.BestIndividual())
	}
}

func TestABC_EmployeeStage(t *testing.T) {
	abc := NewABC(10, 3, Boundary{
		Lower: []float64{-10, -10, -10},
		Upper: []float64{10, 10, 10},
	}, 10, 3, func(float64s []float64) float64 {
		return float64s[0] + float64s[1] + float64s[2]
	})

	for i := 0; i < 10; i++ {
		abc.EmployeeStage()
		abc.OnlookerStage()
	}

	//fmt.Println(abc)
}

func TestABC_Run(t *testing.T) {
	SetMinimum(false)
	iterNum := 10
	abc := NewABC(10, 3, Boundary{
		Lower: []float64{-10, -10, -10},
		Upper: []float64{10, 10, 10},
	}, iterNum, 3, func(float64s []float64) float64 {
		return float64s[0] * float64s[1] * float64s[2]
	})

	abc.Run()

	for i := 0; i < iterNum; i++ {
		fmt.Println(i+1, "代:", abc.HistoryBestIndividuals[i], "->", abc.HistoryBestFNC[i])
	}
}