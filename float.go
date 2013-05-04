package matrix

import (
	"bytes"
	"fmt"
	"image"
)

// New return a new FloatMatrix with the given size
func NewFloatMatrix(m, n int) *FloatMatrix {
	elements := make([]float64, m*n)
	return &FloatMatrix{m, n, elements}
}

// NewFromElements return a new FloatMatrix with elemtnts
func NewFloatMatrixFromElements(m, n int, elements []float64) *FloatMatrix {
	return &FloatMatrix{m, n, elements}
}

type FloatMatrix struct {
	m, n     int // row, column
	elements []float64
}

func (mat *FloatMatrix) Size() (int, int) {
	return mat.m, mat.n
}

func (mat *FloatMatrix) Row() int {
	return mat.m
}

func (mat *FloatMatrix) Col() int {
	return mat.n
}

func (mat *FloatMatrix) At(i, j int) float64 {
	return mat.elements[i*mat.n+j]
}

// Set value
func (mat *FloatMatrix) Set(i, j int, v float64) {
	mat.elements[i*mat.n+j] = v
}

// Clear all elements to zero.
func (mat *FloatMatrix) Clear() {
	for i := 0; i < mat.m*mat.n; i++ {
		mat.elements[i] = 0.0
	}
}

func (mat *FloatMatrix) String() string {
	buffer := bytes.NewBufferString("")

	for i := 0; i < mat.m; i++ {
		for j := 0; j < mat.n; j++ {
			fmt.Fprint(buffer, mat.At(i, j), " ")
		}
		fmt.Fprintln(buffer)
	}

	return string(buffer.Bytes())
}

func (mat *FloatMatrix) Elem() []float64 {
	return mat.elements
}

func (mat *FloatMatrix) Find(num float64, p chan image.Point) {
	for i := 0; i < mat.m; i++ {
		for j := 0; j < mat.n; j++ {
			if mat.At(i, j) == num {
				p <- image.Point{i, j}
			}
		}
	}
	return
}

func (a *FloatMatrix) Eq(b *FloatMatrix) bool {
	if a.m != b.m || a.n != b.n {
		return false
	}

	for i := 0; i < a.m; i++ {
		for j := 0; j < a.n; j++ {
			if a.At(i, j) != b.At(i, j) {
				return false
			}
		}
	}
	return true
}

func (mat *FloatMatrix) Transpose() *FloatMatrix {
	out := NewFloatMatrix(mat.n, mat.m)

	for i := 0; i < mat.m; i++ {
		for j := 0; j < mat.n; j++ {
			out.Set(j, i, mat.At(i, j))
		}
	}
	return out
}
