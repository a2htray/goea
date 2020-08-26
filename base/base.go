// Package base provides the basis of evolutionary algorithms
package base

import (
	"github.com/a2htray/goea/math"
	"math/rand"
	"os"
	"sort"
	"time"
)

// rng represents the rand number generator
var (
	rng = rand.New(rand.NewSource(time.Now().Unix()))
)

// CompareWithCalculate ...
func CompareWithCalculate(i1, i2 Individual, minimum bool, fc func([]float64) float64) Individual {
	i1.FitnessValue = fc(i1.Vector)
	i2.FitnessValue = fc(i2.Vector)
	if minimum {
		if i1.FitnessValue > i2.FitnessValue {
			return i2
		} else {
			return i1
		}
	} else {
		if i1.FitnessValue > i2.FitnessValue {
			return i1
		} else {
			return i2
		}
	}
}

// Individual ...
type Individual struct {
	Vector math.Vector
	FitnessValue float64
}

// NewIndividual ...
func NewIndividual(n int, upper Boundary, lower Boundary) Individual {
	diffVector, err := math.VectorSub(upper.Vector(), lower.Vector())
	if err != nil {
		panic(err)
	}
	vector := make(math.Vector, n, n)
	for i := 0; i < n; i++ {
		vector[i] = lower.Vector()[i] + rng.Float64() * diffVector[i]
	}
	return Individual{
		Vector: vector,
		FitnessValue: 0,
	}
}

// Copy
func (i Individual) Copy() Individual {
	return Individual{
		Vector: i.Vector.Copy(),
		FitnessValue: i.FitnessValue,
	}
}

// CompareTo ...
func (i Individual) CompareTo(i1 Individual, minimum bool, fc func([]float64) float64) bool {
	if i.FitnessValue == 0 {
		i.FitnessValue = fc(i.Vector)
	}
	if i1.FitnessValue == 0 {
		i1.FitnessValue = fc(i1.Vector)
	}

	if minimum {
		return i.FitnessValue < i1.FitnessValue
	} else {
		return i.FitnessValue > i1.FitnessValue
	}
}

// Population ...
type Population []Individual

// Copy ...
func (p Population) Copy() Population {
	ret := make([]Individual, len(p), len(p))
	for i, individual := range p {
		ret[i] = individual.Copy()
	}
	return ret
}

// EAModel ...
type EAModel struct {
	M, N int
	Population Population
	FC func([]float64) float64
	Limit
	IterNum int
	Minimum bool
	GlobalBestIndividual Individual
	HistoryBestIndividuals []Individual
	LogFlag bool
	Log *os.File
}

func (e *EAModel) OpenLog(wr *os.File) {
	e.LogFlag = true
	e.Log = wr
}


// ComputeCurrentFitnessValues ...
func (e *EAModel) ComputeCurrentFitnessValues()  {
	for _, individual := range e.Population {
		individual.FitnessValue = e.FC(individual.Vector)
	}
}

// CurrentBestIndividual ...
func (e *EAModel) CurrentBestIndividual() Individual {
	copyPopulation := e.Population.Copy()
	sort.Slice(copyPopulation, func(i, j int) bool {
		if e.Minimum {
			return copyPopulation[i].FitnessValue < copyPopulation[j].FitnessValue
		} else {
			return copyPopulation[i].FitnessValue > copyPopulation[j].FitnessValue
		}
	})

	return copyPopulation[0]
}

// initPopulation ...
func (e *EAModel) initPopulation() {
	e.Population = make([]Individual, e.M, e.M)
	diffVector, err := math.VectorSub(e.Upper.Vector(), e.Lower.Vector())

	if err != nil {
		panic(err)
	}

	for i := 0; i < e.M; i++ {
		vector := make([]float64, e.N, e.N)
		for j := 0; j < e.N; j++ {
			vector[j] = diffVector[j]*rng.Float64() + e.Lower[j]
		}
		e.Population[i] = Individual{
			Vector: vector,
			FitnessValue: e.FC(vector),
		}
	}
}

// NewEAModel ...
func NewEAModel(m, n int, iterNum int, minimum bool, limit Limit, fc func([]float64) float64) *EAModel {
	model := &EAModel{
		M: m,
		N: n,
		IterNum: iterNum,
		FC: fc,
		Limit: limit,
		Minimum: minimum,
		HistoryBestIndividuals: make([]Individual, iterNum, iterNum),
		LogFlag: false,
	}
	model.initPopulation()
	model.GlobalBestIndividual = model.CurrentBestIndividual()
	model.HistoryBestIndividuals[0] = model.GlobalBestIndividual.Copy()

	return model
}



