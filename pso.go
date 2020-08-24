package goea

import (
	"errors"
	"fmt"
)

// PSOVelocityBoundary PSO 中粒子的速度限制
type PSOVelocityBoundary struct {
	Upper []float64
	Lower []float64
}

// PSOConfig PSO 配置结构
type PSOConfig struct {
	C1, C2           float64
	VelocityBoundary PSOVelocityBoundary
}

// PSO 算法
type PSO struct {
	*eaModel
	PSOConfig
	velocities []float64
}

// GetVelocity 得到第 i 个粒子的速度
func (p *PSO) GetVelocity(i int) []float64 {
	return p.velocities[i*p.N:i*p.N+p.N]
}

// SetVelocity 设置第 i 个粒子的速度
func (p *PSO) SetVelocity(i int, velocity []float64)  {
	for j, v := range velocity {
		p.velocities[i*p.N+j] = v
	}
}

// Run 执行 PSO
func (p *PSO) Run()  {
	for t := 0; t < p.IterNum; t++ {
		// currentBest 是当前最优的解
		currentBest, _ := p.BestIndividual()
		// globalBest 是历史最优的解
		historyBest := p.GetHistoryBestIndividual(t)
		fmt.Println(historyBest)
		for i := 0; i < p.M; i++ {
			for j := 0; j < p.N; j++ {
				fmt.Println(j)
				p.velocities[i*p.N+j] = p.velocities[i*p.N+j] +
					p.C1 * rng.Float64() * (currentBest[j] - p.Population[i][j]) +
					p.C2 * rng.Float64() * (historyBest[j] -
						p.Population[i][j])
			}
		}
		p.HistoryBestIndividuals[t], p.HistoryBestFNC[t] = p.BestIndividual()
		fmt.Println(p.velocities)
	}
}

// DefaultPSOConfig 生成默认的 PSO 配置
func DefaultPSOConfig() PSOConfig {
	return PSOConfig{
		C1: 0.5,
		C2: 0.9,
		VelocityBoundary: PSOVelocityBoundary{
			Upper: []float64{-5},
			Lower: []float64{5},
		},
	}
}

var (
	errorPSORangeNotEqual = errors.New("the range is not equal")
)

// NewPSO 生成新的 PSO 算法模型
func NewPSO(m, n int, boundary Boundary, iterNum int, fc func([]float64) float64, config PSOConfig) *PSO {
	pso := new(PSO)
	pso.eaModel = newEAModel(m, n, boundary, iterNum, fc)
	pso.PSOConfig = config

	if len(pso.VelocityBoundary.Lower) == 1 {
		v := pso.VelocityBoundary.Lower[0]
		for i := 1; i < n; i++ {
			pso.VelocityBoundary.Lower = append(pso.VelocityBoundary.Lower, v)
		}
	}

	if len(pso.VelocityBoundary.Upper) == 1 {
		v := pso.VelocityBoundary.Upper[0]
		for i := 1; i < n; i++ {
			pso.VelocityBoundary.Upper = append(pso.VelocityBoundary.Upper, v)
		}
	}

	//fmt.Println(len(pso.VelocityBoundary.Lower), len(pso.VelocityBoundary.Upper))
	if len(pso.VelocityBoundary.Lower) != n || len(pso.VelocityBoundary.Upper) != n {
		panic(errorPSORangeNotEqual)
	}

	pso.AssignBestIndividual()

	pso.velocities = make([]float64, m * n)
	for i := 0; i < m; i++ {
		vector := newVector(n, pso.VelocityBoundary.Upper, pso.VelocityBoundary.Lower)
		for j := 0; j < n; j++ {
			pso.velocities[i*n+j] = vector[j]
		}
	}

	return pso
}