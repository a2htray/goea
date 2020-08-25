// Package base provides the basis of evolutionary algorithms
package base

import (
	"github.com/a2htray/goea/math"
	"math/rand"
	"sort"
	"time"
)

// rng represents the rand number generator
var (
	rng = rand.New(rand.NewSource(time.Now().Unix()))
)

// Individual ...
type Individual struct {
	Vector math.Vector
	FitnessValue float64
}

// Copy
func (i Individual) Copy() Individual {
	return Individual{
		Vector: i.Vector.Copy(),
		FitnessValue: i.FitnessValue,
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

// eaModel ...
type eaModel struct {
	M, N int
	Population Population
	FC func([]float64) float64
	Limit
	IterNum int
	Minimum bool
	GlobalBestIndividual Individual
	HistoryBestIndividuals []Individual
}

// ComputeCurrentFitnessValues ...
func (e *eaModel) ComputeCurrentFitnessValues()  {
	for _, individual := range e.Population {
		individual.FitnessValue = e.FC(individual.Vector)
	}
}

// CurrentBestIndividual
func (e *eaModel) CurrentBestIndividual() Individual {
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
func (e *eaModel) initPopulation() {
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
func NewEAModel(m, n int, iterNum int, minimum bool, limit Limit, fc func([]float64) float64) *eaModel {
	model := &eaModel{
		M: m,
		N: n,
		IterNum: iterNum,
		FC: fc,
		Limit: limit,
		Minimum: minimum,
		HistoryBestIndividuals: make([]Individual, iterNum, iterNum),
	}
	model.initPopulation()
	model.GlobalBestIndividual = model.CurrentBestIndividual()
	model.HistoryBestIndividuals[0] = model.GlobalBestIndividual.Copy()

	return model
}



