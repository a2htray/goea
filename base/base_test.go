package base

import (
	"fmt"
	"github.com/a2htray/goea"
	"sort"
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

func TestA(t *testing.T) {
	type Item struct {
		Value int
		Weight int
	}

	items := []Item{
		{Value: 1,Weight: 1},
		{Value: 3,Weight: 5},
		{Value: 2,Weight: 6},
		{Value: 5,Weight: 4},
		{Value: 6,Weight: 3},
		{Value: 0,Weight: 2},
	}

	// 按 Value 升序
	sort.Slice(items, func(i, j int) bool {
		return items[i].Value < items[j].Value
	})
	fmt.Println("按 Value 升序")
	for _, item := range items {
		fmt.Println("Value: ", item.Value, "Weight", item.Weight)
	}

	// 按 Weight 降序
	sort.Slice(items, func(i, j int) bool {
		return items[i].Weight > items[j].Weight
	})
	fmt.Println("按 Weight 降序")
	for _, item := range items {
		fmt.Println("Value: ", item.Value, "Weight", item.Weight)
	}
}