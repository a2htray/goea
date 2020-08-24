package goea

// Vector 个体
type Vector []float64

// N 返回向量分量个数
func (v Vector) N() (n int) {
	n = len(v)
	return
}

// Mat 1xN => Nx1
func (v Vector) Mat() Mat {
	n := v.N()
	mat := make([]Vector, n)
	for i := 0; i < n; i++ {
		mat[i] = Vector{v[i]}
	}
	return mat
}

// newVector 生成向量
func newVector(n int, upper []float64, lower []float64) []float64 {
	vector := make([]float64, 0)
	diff := Subtract(upper, lower)
	for i := 0; i < n; i++ {
		vector = append(vector, lower[0] + rng.Float64()*diff[i])
	}
	return vector
}