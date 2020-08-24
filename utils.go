package goea

import (
	"errors"
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
func RandIntRange(min, max int) int {
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
func ShuffleSliceInt(slice []int) []int {
	rng.Shuffle(len(slice), func(i, j int) {
		slice[i], slice[j] = slice[j], slice[i]
	})
	return slice
}

// SumFloat64 求和
func SumFloat64(slice []float64) float64 {
	ret := float64(0)
	for _, v := range slice {
		ret += v
	}
	return ret
}

// DivideSliceFloat64 slice 除以一个除数
func DivideSliceFloat64(slice []float64, divider float64) []float64 {
	ret := make([]float64, len(slice))
	for i, v := range slice {
		ret[i] = v / divider
	}
	return ret
}

// SubtractSliceFloat 相减
func SubtractSliceFloat(slice1 []float64, slice2 []float64) []float64 {
	n1 := len(slice1)
	n2 := len(slice2)

	var n int
	switch n1 > n2 {
	case true:
		n = n2
	default:
		n = n1
	}

	ret := make([]float64, n, n)
	for i := 0; i < n; i++ {
		ret[i] = slice1[i] + slice2[2]
	}
	return ret
}

// AddSliceFloat 相加
func AddSliceFloat(slice1 []float64, slice2 []float64) []float64 {
	n1 := len(slice1)
	n2 := len(slice2)

	var n int
	switch n1 > n2 {
	case true:
		n = n2
	default:
		n = n1
	}

	ret := make([]float64, n, n)
	for i := 0; i < n; i++ {
		ret[i] = slice1[i] + slice2[2]
	}
	return ret
}

var (
	errorSumNotOne = errors.New("the sum is not equal to 1")
	errorRangeNotEqual = errors.New("the index range is not equal in both slice")
)

// RandomChoicePdf 依概率选择
func RandomChoicePdfInt(slice []int, pdf []float64) int {
	if len(pdf) == 0 {
		return slice[0]
	}

	var choices []int
	for i, w := range pdf {
		wi := int(w * 10)
		for j := 0; j < wi; j++ {
			choices = append(choices, i)
		}
	}
	return slice[choices[rng.Int()%len(choices)]]
}

// RemoveSliceInt 从整型 slice 中删除特定元素
// 如果 n = -1，则删除所有相同的元素
func RemoveSliceInt(slice []int, needle int, n int) []int {
	l := len(slice)
	if n <= -1 {
		n = len(slice)
	}

	ret := make([]int, 0)
	counter := 0

	for i, v := range slice {
		if v != needle {
			ret = append(ret, v)
		} else {
			counter++
		}

		if counter >= n {
			if i < l {
				ret = append(ret, slice[i+1:]...)
			}
			break
		}
	}

	return ret
}

// FindNearestInt 返回与浮点数最近的整数
func FindNearestInt(v float64) int {
	return int(v+0.5)
}