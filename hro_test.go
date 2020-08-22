package goea

import (
	"fmt"
	"math"
	"testing"
)

func TestNewHRO(t *testing.T) {
	StartIndex = 1
	hro := NewHRO(8, 2, Boundary{
		Lower: []float64{1, 1},
		Upper: []float64{2, 2},
	}, 400, func(xs []float64) float64 {
		return xs[1] * xs[0]
	})

	fmt.Println(*hro)
}

func TestHRO_Sort(t *testing.T) {
	SetMinimum(false)
	hro := NewHRO(5, 2, Boundary{
		Lower: []float64{1, 1},
		Upper: []float64{2, 2},
	}, 400, func(xs []float64) float64 {
		return xs[1] - xs[0]
	})
	hro.Sort()
	//fmt.Println(hro)
}

func TestHRO_Sort2(t *testing.T) {
	hro := NewHRO(5, 2, Boundary{
		Lower: []float64{1, 1},
		Upper: []float64{2, 2},
	}, 400, func(xs []float64) float64 {
		return math.Pow(xs[1], 2) - xs[0]
	})
	hro.Sort()
	fmt.Println(hro)
}

func TestHybridization(t *testing.T) {
	fc := func(xs []float64) float64 {
		return xs[0] - xs[1]
	}
	hro := NewHRO(9, 3, Boundary{
		Lower: []float64{-10, -10, -10},
		Upper: []float64{10, 10, 10},
	}, 400, fc)
	hro.Sort()
	fmt.Println(hro)
	hro.hybridization()
	fmt.Println(hro)
	fmt.Println(hro.Population.ApplyTo(fc))
}

func TestSelfing(t *testing.T) {
	fc := func(xs []float64) float64 {
		return xs[0] + xs[1] + xs[2]
	}
	hro := NewHRO(9, 3, Boundary{
		Lower: []float64{-100, -100, -100},
		Upper: []float64{100, 100, 100},
	}, 400, fc)
	hro.Sort()
	fmt.Println(hro)
	for i := 0; i < 30000; i++ {
		hro.hybridization()
		hro.selfing()
		hro.calculateFNC()
	}
	hro.Sort()
	fmt.Println(hro)
}

func TestHRO(t *testing.T) {
	SetMinimum(false)
	fc := func(xs []float64) float64 {
		return xs[0] + xs[1] + xs[2]
	}
	hro := NewHRO(9, 3, Boundary{
		Lower: []float64{-100, -100, -100},
		Upper: []float64{100, 100, 100},
	}, 400, fc)
	hro.Sort()
	fmt.Println(hro)
	for i := 0; i < 30000; i++ {
		hro.hybridization()
		hro.selfing()
		hro.calculateFNC()

	}
	hro.Sort()
	fmt.Println(hro)
}

func TestHRO_Run(t *testing.T) {
	fc := func(xs []float64) float64 {
		return xs[0] + xs[1] + xs[2]
	}
	hro := NewHRO(9, 3, Boundary{
		Lower: []float64{-100, -100, -100},
		Upper: []float64{100, 100, 100},
	}, 30000, fc)
	fmt.Println(hro)

	hro.Run()

	fmt.Println(hro)
}

func TestHRO_Run2(t *testing.T) {
	fc := func(xs []float64) float64 {
		return xs[0] + xs[1] + xs[2]
	}
	iterNum := 30000
	hro := NewHRO(9, 3, Boundary{
		Lower: []float64{-100, -100, -100},
		Upper: []float64{100, 100, 100},
	}, iterNum, fc)

	hro.Run()

	for i := 0; i < iterNum; i++ {
		fmt.Println(i+1, "代:", hro.HistoryBestIndividuals[i], "->", hro.HistoryBestFNC[i])
	}
}
