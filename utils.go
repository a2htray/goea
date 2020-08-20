package goea

import (
	"errors"
	"math/rand"
)

var (
	errorRowNumNotEqual = errors.New("the number of row is not equal")
)

// Subtract 相减
func Subtract(a, b []float64) (result []float64) {
	for i := 0; i < len(a); i++ {
		result = append(result, a[i]-b[i])
	}
	return
}

// VectorExpand 向量相连
func VectorExpand(v1, v2 Vector) Vector {
	vec := make([]float64, 0, len(v1)+len(v2))
	vec = append(vec, v1...)
	vec = append(vec, v2...)
	return vec
}

// ExpandX 两个矩阵水平相连
func MatExpandX(mat1, mat2 Mat) Mat {
	m1, _ := mat1.MN()
	m2, _ := mat2.MN()

	if m1 != m2 {
		panic(errorRowNumNotEqual)
	}

	mat := make([]Vector, m1)

	for i := 0; i < m1; i++ {
		mat[i] = VectorExpand(mat1[i], mat2[i])
	}

	return mat
}

var (
	errorMaxMin = errors.New("max must be larger than min")
)

// RandIntRange 返回指定区间内的随机整数
func RandIntRange(min, max int, rng *rand.Rand) int {
	if max < min {
		panic(errorMaxMin)
	}
	return rng.Intn(max-min) + min
}

// RangeInt 实现 python 的 range 函数
func RangeInt(min, max int) []int {
	if min > max {
		panic(errorMaxMin)
	}

	ret := make([]int, 0)
	for ; min < max; min++ {
		ret = append(ret, min)
	}
	return ret
}

// ShuffleSliceInt 打乱 int slice 的排序
func ShuffleSliceInt(slice []int, rng *rand.Rand) []int {
	rng.Shuffle(len(slice), func(i, j int) {
		slice[i], slice[j] = slice[j], slice[i]
	})

	return slice
}
