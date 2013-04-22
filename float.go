package matrix

import (
	"bytes"
	"fmt"
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

func (mat *FloatMatrix) Size() (m, n int) {
	return m, n
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
