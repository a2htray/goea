package goea

import "errors"

// 与上下限相关的错误
var (
	errorBoundaryNumIncompatible = errors.New("the numbers of both Upper and Lower are incompatible")
	errorBoundaryComponentValue  = errors.New("the values of components in Upper and Lower are not proper")
	errorBoundaryNumNotThanTo    = errors.New("the given number is larger then the number of components")
)

// 个体生成策略常量
const (
	GenerateStrategyRandom = iota
	GenerateStrategyMedium
	GenerateStrategyBoundary
)

var (
	GenerateStrategySingleFunc = generateStrategyBoundarySingle
	GenerateStrategyFunc = generateStrategyBoundary
)

// SetGenerateStrategy 设置边界策略
func SetGenerateStrategy(strategy int) {
	switch strategy {
	case GenerateStrategyBoundary:
		GenerateStrategySingleFunc = generateStrategyBoundarySingle
		GenerateStrategyFunc = generateStrategyBoundary
	case GenerateStrategyRandom:
		GenerateStrategySingleFunc = generateStrategyRandomSingle
		GenerateStrategyFunc = generateStrategyRandom
	case GenerateStrategyMedium:
		GenerateStrategySingleFunc = generateStrategyMediumSingle
		GenerateStrategyFunc = generateStrategyMedium
	}
}

// generateStrategyBoundarySingle 对单个值进行判断，取其上限或下限
func generateStrategyBoundarySingle(v, upper, lower float64) float64 {
	if v > upper {
		return upper
	}
	if v < lower {
		return lower
	}
	return v
}

// generateStrategyBoundary 按取上下限来解决越界问题，从而生成新个体
func generateStrategyBoundary(individual Individual, boundary Boundary) Individual {
	for i := 0; i < len(individual); i++ {
		individual[i] = generateStrategyBoundarySingle(individual[i], boundary.Upper[i], boundary.Lower[i])
	}
	return individual
}

// generateStrategyRandomSingle 以随机的方式生成个体分量
func generateStrategyRandomSingle(_, upper, lower float64) float64 {
	return rng.Float64() * (upper - lower) + lower
}

// generateStrategyRandom 以随机的方式生成个体
func generateStrategyRandom(individual Individual, boundary Boundary) Individual {
	return NewIndividual(len(individual), boundary)
}

// generateStrategyMediumSingle 以取平均值的方式生成个体
func generateStrategyMediumSingle(_, upper, lower float64) float64 {
	return (upper+lower)/2
}

// generateStrategyMedium 以取平均值的方式生成个体
func generateStrategyMedium(_ Individual, boundary Boundary) Individual {
	return  DivideSliceFloat64(Subtract(boundary.Upper, boundary.Lower), 2)
}