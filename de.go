package goea

import (
	"errors"
)

// DEConfig DE 的超参配置
type DEConfig struct {
	F float64
	CR float64
}

// DefaultDEConfig 得到一个默认的 DE 算法的配置
func DefaultDEConfig() DEConfig {
	return DEConfig{
		F: 0.8,
		CR: 0.9,
	}
}

var (
	errorSimpleM4 = errors.New("the number of individuals must not be less than 4")
)

// SimpleDE simple DE
type SimpleDE struct {
	*eaModel
	DEConfig
}

// Run ...
func (s *SimpleDE) Run() {
	for i := 0; i < s.IterNum; i++ {
		s.Mutation()
		s.HistoryBestIndividuals[i], s.HistoryBestFNC[i] = s.BestIndividual()
	}
}

// Mutation 变异操作
func (s *SimpleDE) Mutation() {
	for i, individual := range s.Population {
		shuffled := ShuffleSliceInt(RemoveSliceInt(RangeInt(0, s.M), i, 1))
		i1, i2, i3 := shuffled[0], shuffled[1], shuffled[2]

		//oldIndividual := individual.Copy()
		//fmt.Println("old: ", oldIndividual, "--fitness:", oldIndividual.ApplyTo(s.FC))

		newIndividual := s.Population[i1].AddIndividual(s.Population[i2].SubIndividual(s.Population[i3])).MultipleIndividual(s.DEConfig.F)
		//fmt.Println("after mutation: ",newIndividual, "--fitness:", newIndividual.ApplyTo(s.FC))

		// Crossover 操作
		newIndividual = s.Crossover(individual, newIndividual)
		//fmt.Println("after crossover: ", newIndividual, "--fitness:", newIndividual.ApplyTo(s.FC))

		// Selection 选择操作
		s.Population[i] = compare(individual, newIndividual, s.FC)
		//fmt.Println("after selection: ", s.Population[i], "--fitness:", s.Population[i].ApplyTo(s.FC))

		//fmt.Println()
	}
}

// Crossover 交叉操作
func (s *SimpleDE) Crossover(older, newer Individual) Individual {
	for i := 0; i < s.N; i++ {
		r := rng.Float64()

		if !(r <= s.DEConfig.CR || RandIntRange(0, s.N) != i) {
			newer[i] = older[i]
		}
	}
	return newer
}

func NewSimpleDE(m, n int, boundary Boundary, iterNum int, fc func([]float64) float64, config DEConfig) *SimpleDE {
	// 在 simple DE 中，个体的个数必须大于等于 4
	if m < 4 {
		panic(errorSimpleM4)
	}
	de := new(SimpleDE)
	de.eaModel = newEAModel(m, n, boundary, iterNum, fc)
	de.DEConfig = config

	de.HistoryBestIndividuals[0], de.HistoryBestFNC[0] = de.BestIndividual()

	return de
}

