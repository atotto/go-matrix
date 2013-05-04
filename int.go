package matrix

import (
	"bytes"
	"fmt"

	"reflect"
)

type Mat struct {
	m, n int
	elem []interface{}
	kind reflect.Kind
}

type Elem []interface{}

func NewMat(m, n int, elem []interface{}, kind reflect.Kind) *Mat {
	return &Mat{m, n, elem, kind}
}

func (mat *Mat) Add(matB *Mat) {
	switch mat.kind {
	case reflect.Int:
		for i := 0; i < mat.m*mat.n; i++ {
			r := mat.elem[i].(int) + matB.elem[i].(int)
			mat.elem[i] = r
		}
		break
	case reflect.Float64:
		for i := 0; i < mat.m*mat.n; i++ {
			r := mat.elem[i].(float64) + matB.elem[i].(float64)
			mat.elem[i] = r
		}
		break
	default:
		panic("Does not support:" + mat.kind.String())
	}
}

func (mat *Mat) At(i, j int) interface{} {
	return mat.elem[i*mat.n+j]
}

func (mat *Mat) String() string {
	buffer := bytes.NewBufferString(mat.kind.String())
	fmt.Fprintln(buffer)

	for i := 0; i < mat.m; i++ {
		for j := 0; j < mat.n; j++ {
			fmt.Fprint(buffer, mat.At(i, j), " ")
		}
		fmt.Fprintln(buffer)
	}

	return string(buffer.Bytes())
}
