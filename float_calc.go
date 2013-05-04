package matrix

func (matA *FloatMatrix) Mul(matB *FloatMatrix) *FloatMatrix {
	if matA.n != matB.m {
		panic("The left matrix must be the same as the number of rows of the right matrix")
	}

	out := NewFloatMatrix(matA.m, matB.n)

	for i := 0; i < out.m; i++ {
		for j := 0; j < out.n; j++ {
			for n := 0; n < matA.n; n++ {
				out.Set(i, j, out.At(i, j)+matA.At(i, n)*matB.At(n, j))
			}
		}
	}
	return out
}
