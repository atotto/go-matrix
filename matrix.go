package matrix

import (
	"bytes"
	"fmt"
	"image"
)

// NewMatrix return a new Matrix with the given size
func NewMatrix[T int | float64](m, n int) *Matrix[T] {
	elements := make([]T, m*n)
	return &Matrix[T]{m, n, elements}
}

// NewMatrixFromElements return a new Matrix with elements
func NewMatrixFromElements[T int | float64](m, n int, elements []T) *Matrix[T] {
	return &Matrix[T]{m, n, elements}
}

type Matrix[T int | float64] struct {
	m, n     int // row, column
	elements []T
}

func (m *Matrix[T]) Size() (int, int) {
	return m.m, m.n
}

func (m *Matrix[T]) Row() int {
	return m.m
}

func (m *Matrix[T]) Col() int {
	return m.n
}

func (m *Matrix[T]) At(i, j int) T {
	return m.elements[i*m.n+j]
}

// Set value
func (m *Matrix[T]) Set(i, j int, v T) {
	m.elements[i*m.n+j] = v
}

// Clear all elements to zero.
func (m *Matrix[T]) Clear() {
	for i := 0; i < m.m*m.n; i++ {
		m.elements[i] = 0.0
	}
}

func (m *Matrix[T]) String() string {
	buffer := bytes.NewBufferString("")

	for i := 0; i < m.m; i++ {
		for j := 0; j < m.n; j++ {
			fmt.Fprint(buffer, m.At(i, j), " ")
		}
		fmt.Fprintln(buffer)
	}

	return string(buffer.Bytes())
}

func (m *Matrix[T]) Elem() []T {
	return m.elements
}

func (m *Matrix[T]) Find(num T, p chan image.Point) {
	for i := 0; i < m.m; i++ {
		for j := 0; j < m.n; j++ {
			if m.At(i, j) == num {
				p <- image.Point{i, j}
			}
		}
	}
	return
}

func (m *Matrix[T]) Eq(b *Matrix[T]) bool {
	if m.m != b.m || m.n != b.n {
		return false
	}

	for i := 0; i < m.m; i++ {
		for j := 0; j < m.n; j++ {
			if m.At(i, j) != b.At(i, j) {
				return false
			}
		}
	}
	return true
}

func (m *Matrix[T]) Transpose() *Matrix[T] {
	out := NewMatrix[T](m.n, m.m)

	for i := 0; i < m.m; i++ {
		for j := 0; j < m.n; j++ {
			out.Set(j, i, m.At(i, j))
		}
	}
	return out
}
