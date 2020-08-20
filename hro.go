package goea

import (
	"sort"
	"strconv"
	"strings"
)

var hroLineNum = 3
var SelfingMaxNum = 10

// HRO 三系杂交水稻算法
// 参考文献 1
type HRO struct {
	Population
	M, N int
	Boundary
	IterNum int
	// FC 目标函数
	FC func([]float64) float64
	// FNC 适应值集合
	FNC      []float64
	// 各系植株的个数
	LineSize int
	// 每一代最优的适应值
	perFNC []float64
	// 每一代最优的个体
	perIndividuals []Individual
}

// calculateFNC 计算适应值
func (h *HRO) calculateFNC() {
	h.FNC = h.ApplyTo(h.FC)
}

// String 算法字符串打印
func (h HRO) String() string {
	strs := []string{h.Population.String()}

	for i, fitness := range h.FNC {
		strs = append(strs, strconv.Itoa(i+StartIndex)+" "+strconv.FormatFloat(fitness, 'f', -1, 64))
	}

	return strings.Join(strs, "\n")
}

// hybridization 杂交
func (h *HRO) hybridization() {
	for _, maintainer := range h.Population[0:h.LineSize] {
		// 随机选择的不育系植株下标
		sterileIndex := RandIntRange(h.LineSize*2, h.M, rng)
		newSterile := make(Individual, h.N, h.N)
		for i := 0; i < h.N; i++ {
			// r1, r2 [0.0, 1.0)
			r1, r2 := rng.Float64(), rng.Float64()
			newGene := (r1*maintainer[i] + r2*h.Population[sterileIndex][i]) / (r1 + r2)
			newSterile[i] = newGene
		}
		h.Population[sterileIndex] = compare(h.Population[sterileIndex], newSterile, h.FC)
	}
}

// selfing 自交
func (h *HRO) selfing() {
	selfingStart := h.LineSize + 1
	selfingEnd := h.LineSize * 2

	for i, selfinger := range h.Population[selfingStart:selfingEnd] {
		// 当前恢复系中最优的个体
		best := h.Population[h.LineSize]
		currentIndex := selfingStart + i
		isReplaced := false
		for selfingNum := 0; selfingNum < SelfingMaxNum; selfingNum++ {
			// 第一步是为了算出与自交植株不同植株的下标
			selectedIndex := currentIndex
			for _, index := range ShuffleSliceInt(RangeInt(selfingStart, selfingEnd), rng) {
				if index != selectedIndex {
					selectedIndex = index
				}
			}

			// 随机到基因下标
			geneIndex := RandIntRange(0, h.N, rng)
			r := rng.Float64()

			newGene := r*(best[geneIndex]-h.Population[selectedIndex][geneIndex]) + selfinger[geneIndex]
			newSelfinger := append(selfinger[:geneIndex], newGene)
			newSelfinger = append(newSelfinger, selfinger[geneIndex+1:]...)

			if Individual(newSelfinger).Compare(selfinger, h.FC) {
				isReplaced = true
				h.Population[currentIndex] = Individual(newSelfinger)
				// 如果生成的植株优于当前最优，则两者交换位置
				if h.Population[currentIndex].Compare(best, h.FC) {
					h.Population[h.LineSize], h.Population[currentIndex] = h.Population[currentIndex], h.Population[h.LineSize]
				}
				break
			}
		}

		if !isReplaced {
			h.Population[currentIndex] = h.renewal()
			if h.Population[currentIndex].Compare(best, h.FC) {
				h.Population[h.LineSize], h.Population[currentIndex] = h.Population[currentIndex], h.Population[h.LineSize]
			}
		}
	}
}

func (h *HRO) renewal() Individual {
	return NewIndividual(h.N, h.Boundary)
}

// Sort 进行排序操作
func (h *HRO) Sort() {
	mat := MatExpandX(h.Population.mat(), Vector(h.FNC).Mat())
	_, n := mat.MN()

	sort.Slice(mat, func(i, j int) bool {
		if Minimum {
			return mat[i][n-1] < mat[j][n-1]
		} else {
			return mat[i][n-1] > mat[j][n-1]
		}
	})

	for i, vector := range mat.Cut(0, -1, 0, n-1) {
		h.Population[i] = Individual(vector)
	}

	h.FNC = mat.Cut(0, -1, n-1, -1).Flat()
}

func (h *HRO) Run() {
	for i := 0; i < h.IterNum; i++ {
		h.hybridization()
		h.selfing()
		h.calculateFNC()
		h.Sort()
		h.perFNC[i] = h.FNC[0]
		h.perIndividuals[i] = h.Population[0]
	}
}

// NewHRO 生成 HRO 算法
func NewHRO(m, n int, boundary Boundary, iterNum int, fc func([]float64) float64) (hro *HRO) {
	hro = new(HRO)
	hro.Population = initPopulation(m, n, boundary)
	hro.M = m
	hro.N = n
	hro.Boundary = boundary
	hro.IterNum = iterNum
	hro.FC = fc
	hro.LineSize = m / hroLineNum
	hro.calculateFNC()
	hro.Sort()
	hro.perFNC = make([]float64, iterNum)
	hro.perIndividuals = make([]Individual, iterNum)

	// 最大自交次数不应超过基因的个数
	if SelfingMaxNum > hro.N {
		SelfingMaxNum = hro.N
	}

	return
}
