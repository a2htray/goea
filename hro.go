package goea

import (
	"fmt"
	"github.com/a2htray/goea/base"
	"github.com/a2htray/goea/math"
	"sort"
	"strings"
)

var hroLineNum = 3
var SelfingMaxNum = 10

// HRO 三系杂交水稻算法
// 参考文献 1
type HRO struct {
	*base.EAModel
	// 各系植株的个数
	LineSize int
}

// String 算法字符串打印
func (h HRO) String() string {
	strs := []string{"a", "b"}

	return strings.Join(strs, "\n")
}

// hybridization 杂交
func (h *HRO) hybridization(iter int) {
	if h.LogFlag {
		h.Log.WriteString(fmt.Sprintf("第%d次迭代 hybridization 阶段:\n", iter+1))
	}
	for _, maintainer := range h.Population[0:h.LineSize] {

		// 随机选择的不育系植株下标
		sterileIndex := RandIntRange(h.LineSize*2, h.M)
		if h.LogFlag {
			h.Log.WriteString(fmt.Sprintf("选择保持系植株:(%v,%f)\n", maintainer.Vector, maintainer.FitnessValue))
			h.Log.WriteString(fmt.Sprintf("随机选择的不育系植株:(%v,%f)\n", h.Population[sterileIndex].Vector, h.Population[sterileIndex].FitnessValue))
		}
		newSterile := base.Individual{
			Vector: make(math.Vector, h.N, h.N),
		}
		for i := 0; i < h.N; i++ {
			r1, r2 := rng.Float64(), rng.Float64()
			top := r1*maintainer.Vector[i] + r2*h.Population[sterileIndex].Vector[i]
			bottom := r1 + r2
			newGene :=  top/bottom
			newSterile.Vector[i] = newGene
		}
		newSterile.FitnessValue = h.FC(newSterile.Vector)
		if h.LogFlag {
			h.Log.WriteString(fmt.Sprintf("杂交后新植株:(%v,%f)\n", newSterile.Vector, newSterile.FitnessValue))
		}

		h.Population[sterileIndex] = base.CompareWithCalculate(h.Population[sterileIndex], newSterile, h.Minimum, h.FC)
		if h.LogFlag {
			h.Log.WriteString(fmt.Sprintf("比较后更新的新植株:(%v,%f)\n", h.Population[sterileIndex].Vector, h.Population[sterileIndex].FitnessValue))
		}
	}
}

// selfing 自交
func (h *HRO) selfing(iter int) {
	if h.LogFlag {
		h.Log.WriteString(fmt.Sprintf("第%d次迭代 selfing 阶段:\n", iter+1))
	}
	selfingStart := h.LineSize + 1
	selfingEnd := h.LineSize * 2

	for i, selfinger := range h.Population[selfingStart:selfingEnd] {
		// 当前恢复系中最优的植株
		best := h.Population[h.LineSize]
		if h.LogFlag {
			h.Log.WriteString(fmt.Sprintf("当前恢复系最优植株:(%v,%f)\n", best.Vector, best.FitnessValue))
			h.Log.WriteString(fmt.Sprintf("与当前恢复系最优植株自交的植株:(%v,%f)\n", selfinger.Vector, selfinger.FitnessValue))
		}
		currentIndex := selfingStart + i
		isReplaced := false
		for selfingNum := 0; selfingNum < SelfingMaxNum; selfingNum++ {
			// 第一步是为了算出与自交植株不同植株的下标
			selectedIndex := currentIndex
			for _, index := range ShuffleSliceInt(RangeInt(selfingStart, selfingEnd)) {
				if index != selectedIndex {
					selectedIndex = index
				}
			}

			// 随机到基因下标
			geneIndex := RandIntRange(0, h.N)
			r := rng.Float64()

			newGene := r*(best.Vector[geneIndex]-h.Population[selectedIndex].Vector[geneIndex]) + selfinger.Vector[geneIndex]
			// 使用边界进行判定
			if newGene > h.Upper.Vector()[geneIndex] {
				newGene = h.Upper.Vector()[geneIndex]
			}
			if newGene < h.Lower.Vector()[geneIndex] {
				newGene = h.Lower.Vector()[geneIndex]
			}

			newVector := append(selfinger.Vector[:geneIndex], newGene)
			newVector = append(newVector, selfinger.Vector[geneIndex+1:]...)

			newIndividual := base.Individual{
				Vector: newVector,
				FitnessValue: h.FC(newVector),
			}

			if h.LogFlag {
				h.Log.WriteString(fmt.Sprintf("自交后的新植株:(%v,%f)\n", newIndividual.Vector, newIndividual.FitnessValue))
			}
			
			if newIndividual.CompareTo(selfinger, h.Minimum, h.FC) {
				isReplaced = true
				h.Population[currentIndex] = newIndividual
				if h.LogFlag {
					h.Log.WriteString(fmt.Sprintf("自交后的新植株优于原植株，发生替换:(%v,%f)\n", h.Population[currentIndex].Vector, h.Population[currentIndex].FitnessValue))
				}
				if h.Population[currentIndex].CompareTo(best, h.Minimum, h.FC) {
					h.Population[h.LineSize], h.Population[currentIndex] = h.Population[currentIndex], h.Population[h.LineSize]
					if h.LogFlag {
						h.Log.WriteString(fmt.Sprintf("自交后的新植株优于原最优不育系植株，发生替换:\n"))
						h.Log.WriteString(fmt.Sprintf("当前最优不育系植株:(%v,%f)", h.Population[h.LineSize].Vector, h.Population[h.LineSize].FitnessValue))
					}
				}
			}
		}

		if !isReplaced {
			h.Population[currentIndex] = h.renewal(iter)
			if h.LogFlag {
				h.Log.WriteString(fmt.Sprintf("自交后的新植株次于原植株，发生 renewal:(%v,%f)\n", h.Population[currentIndex].Vector, h.Population[currentIndex].FitnessValue))
			}
			if h.Population[currentIndex].CompareTo(best, h.Minimum, h.FC) {
				h.Population[h.LineSize], h.Population[currentIndex] = h.Population[currentIndex], h.Population[h.LineSize]
				if h.LogFlag {
					h.Log.WriteString(fmt.Sprintf("renewal 后的新植株优于原最优不育系植株，发生替换:\n"))
					h.Log.WriteString(fmt.Sprintf("当前最优不育系植株:(%v,%f)", h.Population[h.LineSize].Vector, h.Population[h.LineSize].FitnessValue))
				}
			}
		}
	}
}

func (h *HRO) renewal(iter int) base.Individual {
	ret := base.NewIndividual(h.N, h.Upper, h.Lower)
	ret.FitnessValue = h.FC(ret.Vector)
	return ret
}

// Sort 进行排序操作
func (h *HRO) Sort() {
	sort.Slice(h.Population, func(i, j int) bool {
		if h.Minimum {
			return h.Population[i].FitnessValue < h.Population[j].FitnessValue
		} else {
			return h.Population[i].FitnessValue > h.Population[j].FitnessValue
		}
	})
}

func (h *HRO) Run() {
	if h.LogFlag {
		_, err := h.Log.WriteString("初始状态:\n")
		if err != nil {
			panic(err)
		}
		for i, individual := range h.Population {
			var tag string
			switch i / h.LineSize {
			case 0:
				tag = "保持系"
			case 1:
				tag = "恢复系"
			default:
				tag = "不育系"

			}
			h.Log.WriteString(fmt.Sprintf("%s:%v,%f\n", tag, individual.Vector, individual.FitnessValue))
		}
		h.Log.WriteString("初始最优解为:\n")
		h.Log.WriteString(fmt.Sprintf("%v,%f\n", h.Population[0].Vector, h.Population[0].FitnessValue))
	}

	for i := 0; i < h.IterNum; i++ {
		h.hybridization(i)
		h.selfing(i)
		h.Sort()
		h.HistoryBestIndividuals[i] = h.Population[0]

		if h.LogFlag {
			h.Log.WriteString(fmt.Sprintf("第%d次迭代后，种群状态:\n", i+1))
			for i, individual := range h.Population {
				var tag string
				switch i / h.LineSize {
				case 0:
					tag = "保持系"
				case 1:
					tag = "恢复系"
				default:
					tag = "不育系"
				}
				h.Log.WriteString(fmt.Sprintf("%s:%v,%f\n", tag, individual.Vector, individual.FitnessValue))
			}
		}
	}
}

// NewHRO 生成 HRO 算法
func NewHRO(m, n int, iterNum int, minimum bool, limit base.Limit, fc func([]float64) float64) (hro *HRO) {
	hro = new(HRO)
	hro.EAModel = base.NewEAModel(m, n, iterNum, minimum, limit, fc)
	hro.LineSize = m / hroLineNum
	hro.Sort()

	// 最大自交次数不应超过基因的个数
	if SelfingMaxNum > hro.N {
		SelfingMaxNum = hro.N
	}

	return
}
