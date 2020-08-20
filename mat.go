package goea

// Mat 矩阵
type Mat []Vector

func (mat *Mat) SetRow(i int, vector Vector) {
	(*mat)[i] = vector
}

// MN 返回矩阵行列数
func (mat Mat) MN() (m, n int) {
	m = len(mat)
	n = len(mat[0])
	return
}

// Cut 取矩阵中的子矩阵
func (mat Mat) Cut(mStart, mEnd, nStart, nEnd int) Mat {
	m, n := mat.MN()

	if mEnd == -1 {
		mEnd = m
	}
	if nEnd == -1 {
		nEnd = n
	}

	newMat := make([]Vector, mEnd-mStart)
	vs := make([]float64, 0, (mEnd-mStart)*(nEnd-nStart))

	for i := mStart; i < mEnd; i++ {
		for j := nStart; j < nEnd; j++ {
			vs = append(vs, mat[i][j])
		}
	}
	for i := range newMat {
		newMat[i], vs = vs[:nEnd-nStart], vs[nEnd-nStart:]
	}

	return newMat
}

// Flat 将矩阵展开成向量
func (mat Mat) Flat() Vector {
	m, n := mat.MN()
	v := make([]float64, 0, m*n)

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			v = append(v, mat[i][j])
		}
	}

	return v
}
