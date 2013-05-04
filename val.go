package matrix

import (
	"bytes"
	"fmt"
	"reflect"
)

type Mat2 struct {
	m, n int
	elem []reflect.Value
}

func NewMat2(m, n int, elem []reflect.Value) *Mat2 {
	return &Mat2{m, n, elem}
}

func (mat *Mat2) At(i, j int) reflect.Value {
	return mat.elem[i*mat.n+j]
}

func (mat *Mat2) String() string {
	buffer := bytes.NewBufferString("")

	for i := 0; i < mat.m; i++ {
		for j := 0; j < mat.n; j++ {
			fmt.Fprint(buffer, mat.At(i, j), " ")
		}
		fmt.Fprintln(buffer)
	}

	return string(buffer.Bytes())
}
