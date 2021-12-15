package matrix

import (
	"bytes"
	"fmt"
	"image"
)

// New return a new FloatMatrix with the given size
func NewFloatMatrix[T int | float64](m, n int) *FloatMatrix[T] {
	elements := make([]T, m*n)
	return &FloatMatrix[T]{m, n, elements}
}

// NewFromElements return a new FloatMatrix with elemtnts
func NewFloatMatrixFromElements[T int | float64](m, n int, elements []T) *FloatMatrix[T]{
	return &FloatMatrix[T]{m, n, elements}
}

type FloatMatrix[T int | float64] struct {
	m, n     int // row, column
	elements []T
}

func (mat *FloatMatrix[T]) Size() (int, int) {
	return mat.m, mat.n
}

func (mat *FloatMatrix[T]) Row() int {
	return mat.m
}

func (mat *FloatMatrix[T]) Col() int {
	return mat.n
}

func (mat *FloatMatrix[T]) At(i, j int) T {
	return mat.elements[i*mat.n+j]
}

// Set value
func (mat *FloatMatrix[T]) Set(i, j int, v T) {
	mat.elements[i*mat.n+j] = v
}

// Clear all elements to zero.
func (mat *FloatMatrix[T]) Clear() {
	for i := 0; i < mat.m*mat.n; i++ {
		mat.elements[i] = 0.0
	}
}

func (mat *FloatMatrix[T]) String() string {
	buffer := bytes.NewBufferString("")

	for i := 0; i < mat.m; i++ {
		for j := 0; j < mat.n; j++ {
			fmt.Fprint(buffer, mat.At(i, j), " ")
		}
		fmt.Fprintln(buffer)
	}

	return string(buffer.Bytes())
}

func (mat *FloatMatrix[T]) Elem() []T {
	return mat.elements
}

func (mat *FloatMatrix[T]) Find(num T, p chan image.Point) {
	for i := 0; i < mat.m; i++ {
		for j := 0; j < mat.n; j++ {
			if mat.At(i, j) == num {
				p <- image.Point{i, j}
			}
		}
	}
	return
}

func (a *FloatMatrix[T]) Eq(b *FloatMatrix[T]) bool {
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

func (mat *FloatMatrix[T]) Transpose() *FloatMatrix[T] {
	out := NewFloatMatrix[T](mat.n, mat.m)

	for i := 0; i < mat.m; i++ {
		for j := 0; j < mat.n; j++ {
			out.Set(j, i, mat.At(i, j))
		}
	}
	return out
}
