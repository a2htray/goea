package goea

import (
	"strconv"
	"strings"
)

type ABC struct {
	*eaModel
	//
	limit int
	Limits []int
}

// String 字符串打印
func (a ABC) String() string {
	strs := []string{a.Population.String()}

	for i, fitness := range a.CurrentFNC {
		strs = append(
			strs,
			strconv.Itoa(i+StartIndex)+" "+strconv.FormatFloat(fitness, 'f', -1, 64) +
				"-limit- " + strconv.Itoa(a.Limits[i]))
	}

	return strings.Join(strs, "\n")
}

// EmployeeStage 雇佣蜂阶段
func (a *ABC) EmployeeStage() {
	for i := 0; i < a.M; i++ {
		individual := NewIndividual(a.N, a.Boundary)
		if a.Population[i].Compare(individual, a.FC) {
			a.Limits[i] += 1
		} else {
			a.Population[i] = individual
			a.Limits[i] = 0
		}
	}
}

// OnlookerStage 观察蜂阶段
func (a *ABC) OnlookerStage()  {
	selectedRange := RangeInt(0, a.M)
	for i := 0; i < a.M; i++ {
		a.CurrentFNC = a.ApplyTo(a.FC)
		nectar := SumFloat64(a.CurrentFNC)

		selectedIndex := RandomChoicePdfInt(selectedRange, DivideSliceFloat64(a.CurrentFNC, nectar))
		for selectedIndex == i {
			selectedIndex = RandomChoicePdfInt(selectedRange, DivideSliceFloat64(a.CurrentFNC, nectar))
		}
		newIndividual := make(Individual, a.N)
		for j := 0; j < a.N; j++ {
			r := rng.Float64() * 2 - 1
			newIndividual[j] = a.Population[i][j] + r * (a.Population[i][j] - a.Population[selectedIndex][j])
		}

		newIndividual = GenerateStrategyFunc(newIndividual, a.Boundary)

		if a.Population[i].Compare(newIndividual, a.FC) {
			a.Limits[i] += 1
		} else {
			a.Population[i] = newIndividual
			a.Limits[i] = 0
		}
	}
}

// ScouterStage 侦察蜂阶段
func (a *ABC) ScouterStage()  {
	for i := 0; i < a.M; i++ {
		if a.Limits[i] >= a.limit {
			a.Population[i] = NewIndividual(a.N, a.Boundary)
			a.Limits[i] = 0
		}
	}
}

func (a *ABC) Run() {
	for i := 0; i < a.IterNum; i++ {
		a.EmployeeStage()
		a.OnlookerStage()
		a.ScouterStage()
		a.HistoryBestIndividuals[i], a.HistoryBestFNC[i] = a.BestIndividual()
	}
}

// NewABC 生成 ABC 算法模型
func NewABC(m, n int, boundary Boundary, iterNum int, limit int, fc func([]float64) float64) *ABC {
	abc := new(ABC)
	abc.eaModel = newEAModel(m, n, boundary, iterNum, fc)
	abc.limit = limit
	abc.Limits = make([]int, m)
	return abc
}