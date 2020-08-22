package goea

import (
	"errors"
	"fmt"
	"math"
	"strconv"
	"strings"
)

const (
	DEModeBest2bin = "best/2/bin"
	DEModeRand1bin = "rand/1/bin"
)

var (
	errorDEModel = errors.New("the passing mode is incorrect")
)

// baseSelector 基向量选择接口
type baseSelector interface {
	Select(population Population, currentIndex int, fc func([]float64) float64) Individual
}

// BestIndividualSelector 当前最优选择器
type BestIndividualSelector struct {}

// Select 选择
func (b *BestIndividualSelector) Select(population Population, _ int, fc func([]float64) float64) Individual {
	p := population.Copy()
	individual, _ := bestIndividual(p, fc)
	return individual
}

type RandIndividualSelector struct {}

// Select 选择
func (b *RandIndividualSelector) Select(population Population, _ int, _ func([]float64) float64) Individual {
	 return population[int(math.Floor(rng.Float64() * float64(population.M())))].Copy()
}

type CurrentIndividualSelector struct {}

// Select 选择
func (b *CurrentIndividualSelector) Select(population Population, currentIndex int, _ func([]float64) float64) Individual {
	return population[currentIndex].Copy()
}

var (
	bSelector = BestIndividualSelector{}
	rSelector = RandIndividualSelector{}
	cSelector = CurrentIndividualSelector{}
)


// 解析 DE 类型
func parseDEMode(mode string) (baseSelector, int, interface{}) {
	modeSegs := strings.Split(mode, "/")

	if len(modeSegs) != 3 {
		panic(errorDEModel)
	}
	fmt.Println(modeSegs)
	var selector baseSelector
	switch modeSegs[0] {
	case "best":
		selector = &bSelector
	case "rand":
		selector = &rSelector
	case "current":
		selector = &cSelector
	}

	diffNum, err := strconv.Atoi(modeSegs[1])
	if err != nil {
		panic(errorDEModel)
	}

	return selector, diffNum, nil

}

// DEConfig DE 的超参配置
type DEConfig struct {
	F float64
	CR float64
	Mode string
	DiffNum int
}

// DefaultDEConfig 得到一个默认的 DE 算法的配置
func DefaultDEConfig() DEConfig {
	return DEConfig{
		F:    0.8,
		CR:   0.9,
		Mode: DEModeRand1bin,
	}
}

var (
	errorSimpleM4 = errors.New("the number of individuals must not be less than 4")
)

// DE simple DE
type DE struct {
	*eaModel
	DEConfig
	selector baseSelector
}

// Run ...
func (d *DE) Run() {
	for i := 0; i < d.IterNum; i++ {
		d.Mutation()
		d.HistoryBestIndividuals[i], d.HistoryBestFNC[i] = d.BestIndividual()
	}
}

// diff 做差
func (d *DE) diff(population Population, indexes... int) Individual {
	p := population.Copy()
	n := len(indexes)
	individual := p[0].SubIndividual(p[1])
	for i := 2; i < n-1; i += 2 {
		individual = individual.AddIndividual(p[0].SubIndividual(p[1]))
	}
	return individual
}

// Mutation 变异操作
func (d *DE) Mutation() {
	for i, individual := range d.Population {
		base := d.selector.Select(d.Population, i, d.FC)

		shuffled := ShuffleSliceInt(RemoveSliceInt(RangeInt(0, d.M), i, d.DiffNum * 2))
		//i2, i3 := shuffled[0], shuffled[1]

		newIndividual := base.AddIndividual(d.diff(d.Population, shuffled...)).MultipleIndividual(d.DEConfig.F)

		// Crossover 操作
		newIndividual = d.Crossover(individual, newIndividual)

		// Selection 选择操作
		d.Population[i] = compare(individual, newIndividual, d.FC)
	}
}

// Crossover 交叉操作
func (d *DE) Crossover(older, newer Individual) Individual {
	for i := 0; i < d.N; i++ {
		r := rng.Float64()

		if !(r <= d.DEConfig.CR || RandIntRange(0, d.N) != i) {
			newer[i] = older[i]
		}
	}
	return newer
}

func NewDE(m, n int, boundary Boundary, iterNum int, fc func([]float64) float64, config DEConfig) *DE {
	// 在 simple DE 中，个体的个数必须大于等于 4
	//if m < 4 {
	//	panic(errorSimpleM4)
	//}
	de := new(DE)
	de.eaModel = newEAModel(m, n, boundary, iterNum, fc)
	de.DEConfig = config

	de.selector, de.DiffNum, _ = parseDEMode(de.DEConfig.Mode)

	fmt.Println(de.selector)

	de.HistoryBestIndividuals[0], de.HistoryBestFNC[0] = de.BestIndividual()

	return de
}

