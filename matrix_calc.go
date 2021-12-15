package matrix

import (
	"fmt"
)

func (m *Matrix[T]) Mul(matB *Matrix[T]) *Matrix[T] {
	if m.n != matB.m {
		panic(fmt.Sprintf("The left matrix (%dx%d) must be the same as the number of rows of the right matrix (%dx%d)", m.m, m.n, matB.m, matB.n))
	}

	out := NewMatrix[T](m.m, matB.n)

	for i := 0; i < out.m; i++ {
		for j := 0; j < out.n; j++ {
			for n := 0; n < m.n; n++ {
				out.Set(i, j, out.At(i, j)+m.At(i, n)*matB.At(n, j))
			}
		}
	}
	return out
}
