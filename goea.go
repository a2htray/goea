package goea

import (
	"errors"
	"fmt"
	"math/rand"
	"sort"
	"strconv"
	"strings"
	"time"
)

var StartIndex = 0
var Minimum = true

func SetMinimum(minimum bool) {
	Minimum = minimum
}

var rng = rand.New(rand.NewSource(time.Now().UnixNano()))

// 与个数相关的错误
var (
	errorNumOfComponents  = errors.New("the number of components must be larger then 0")
	errorNumOfIndividuals = errors.New("the number of individuals must be larger then zero")
)

// Individual 个体
type Individual Vector

// Copy 复制
func (i Individual) Copy() Individual {
	c := make(Individual, len(i))
	copy(c, i)
	return c
}

func (i Individual) MultipleIndividual(value float64) Individual {
	c := i.Copy()
	for index, v := range i {
		c[index] = v * value
	}
	return c
}

// AddIndividual 与个体相加
func (i Individual) AddIndividual(value Individual) Individual {
	c := value.Copy()
	for index, v := range i {
		c[index] = c[index] + v
	}
	return c
}

// SubIndividual 与个体相减
func (i Individual) SubIndividual(value Individual) Individual {
	c := value.Copy()
	for index, v := range i {
		c[index] = v - c[index]
	}
	return c
}

// ApplyTo 将个体应用于特定目标函数，并返回结果
func (i Individual) ApplyTo(fc func([]float64) float64) (fitness float64) {
	fitness = fc(i)
	return
}

// Compare 个体间比较，若优于个体参数，返回 true
func (i Individual) Compare(i1 Individual, fc func([]float64) float64) bool {
	switch i.ApplyTo(fc) < i1.ApplyTo(fc) {
	case true:
		return Minimum
	default:
		return !Minimum
	}
}

// String 字符串输出
func (i Individual) String() string {
	strs := make([]string, 0, len(i))
	for _, c := range i {
		strs = append(strs, strconv.FormatFloat(c, 'f', -1, 64))
	}
	return strings.Join(strs, ",")
}

// Population 种群
type Population []Individual

// Copy
func (p Population) Copy() Population {
	c := make(Population, 0, len(p))
	for _, individual := range p {
		c = append(c, individual.Copy())
	}
	return c
}

// M 返回种群个数
func (p Population) M() int {
	return len(p)
}

// mat 种群变矩阵
func (p Population) mat() Mat {
	mat := make([]Vector, len(p))
	for i, individual := range p {
		mat[i] = Vector(individual)
	}
	return mat
}

// String 种群的字符串输出
func (p Population) String() string {
	strs := make([]string, 0, len(p))
	for i, individual := range p {
		strs = append(strs, strconv.Itoa(i+StartIndex)+" "+individual.String())
	}
	return strings.Join(strs, "\n")
}

// ApplyTo 对整个种群中的各个个体应用于特定目标函数，并返回结果的集合
func (p Population) ApplyTo(fc func([]float64) float64) (fitnessCollection []float64) {
	for _, individual := range p {
		fitnessCollection = append(fitnessCollection, fc(individual))
	}
	return
}

// Boundary 各分量上限及下限
type Boundary struct {
	Upper Vector
	Lower Vector
}

// NN 返回上下限分量个数
func (b Boundary) NN() (un, ln int) {
	un = len(b.Upper)
	ln = len(b.Lower)
	return
}

// CheckSelf 上下限分量自检
func (b Boundary) CheckSelf() {
	un, ln := b.NN()
	if un != ln {
		panic(errorBoundaryNumIncompatible)
	}

	for i := 0; i < len(b.Lower); i++ {
		if b.Lower[i] > b.Upper[i] {
			panic(errorBoundaryComponentValue)
		}
	}
}

type eaModel struct {
	Population
	Boundary
	M, N int
	IterNum int
	// FC 目标函数
	FC func([]float64) float64
	// CurrentFNC 适应值集合
	CurrentFNC []float64
	// HistoryBestFNC 每一代最优的适应值
	HistoryBestFNC []float64
	// HistoryBestIndividuals 每一代最优的个体
	HistoryBestIndividuals []Individual
}

// bestIndividual 求当前种群最好的个体及其适应值
func bestIndividual(population Population, fc func([]float64) float64) (Individual, float64) {
	sort.Slice(population, func(i, j int) bool {
		return population[i].Compare(population[j], fc)
	})

	return population[0].Copy(), population[0].ApplyTo(fc)
}

// BestIndividual 求当前种群最好的个体及其适应值
func (e *eaModel) BestIndividual() (Individual, float64) {
	population := e.Population.Copy()
	fmt.Println(population)
	return bestIndividual(population, e.FC)
}

// calculateFNC 计算适应值
func (e *eaModel) calculateFNC() {
	e.CurrentFNC = e.ApplyTo(e.FC)
}

func newEAModel(m, n int, boundary Boundary, iterNum int, fc func([]float64) float64) *eaModel {
	model := &eaModel{
		Population:             initPopulation(m, n, boundary),
		M:                      m,
		N:                      n,
		Boundary:               boundary,
		IterNum:                iterNum,
		FC:                     fc,
		HistoryBestFNC:         make([]float64, iterNum),
		HistoryBestIndividuals: make([]Individual, iterNum),
	}
	model.calculateFNC()
	return model
}

func initPopulation(m, n int, boundary Boundary) (population Population) {
	if m <= 0 {
		panic(errorNumOfIndividuals)
	}

	if n < 0 {
		panic(errorNumOfComponents)
	}

	boundary.CheckSelf()

	for i := 0; i < m; i++ {
		population = append(population, NewIndividual(n, boundary))
	}
	return
}

// NewIndividual 生成一个新个体
func NewIndividual(n int, boundary Boundary) (individual Individual) {
	un, ln := boundary.NN()
	if n > un || n > ln {
		panic(errorBoundaryNumNotThanTo)
	}

	diff := Subtract(boundary.Upper, boundary.Lower)
	for i := 0; i < n; i++ {
		individual = append(individual, boundary.Lower[i]+rng.Float64()*diff[i])
	}
	return
}

// compare 传递目标函数并比较两个个体，返回较优者
func compare(i1, i2 Individual, fc func([]float64) float64) Individual {
	if i1.Compare(i2, fc) {
		return i1
	} else {
		return i2
	}
}
