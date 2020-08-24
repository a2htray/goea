package goea

import (
	"testing"
)

func TestNewPSO(t *testing.T) {
	iterNum := 100
	pso := NewPSO(6, 3, Boundary{
		Upper: []float64{10, 10, 10},
		Lower: []float64{-10, -10, -10},
	}, iterNum, ObjectSphere, DefaultPSOConfig())

	pso.Run()
}