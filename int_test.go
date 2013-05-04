package matrix

import (
	"reflect"
	"testing"
)

func TestAdd(t *testing.T) {
	a := NewMat(2, 3, []interface{}{2, 1, 2, 4, 5, 6}, reflect.Int)
	b := NewMat(2, 3, []interface{}{3, 6, 4, 2, 6, 7}, reflect.Int)

	a.Add(b)

	println(a.String())
}

func TestAddFloat64(t *testing.T) {
	a := NewMat(2, 3, []interface{}{2.3, 1.2, 2.43, 4.2, 2.5, 0.6}, reflect.Float64)
	b := NewMat(2, 3, []interface{}{2.3, 0.6, 4.2, 0.2, 6.2, 1.17}, reflect.Float64)

	a.Add(b)

	println(a.String())
}

func TestType(t *testing.T) {

}

// func TestAdd2(t *testing.T) {
// 	a := NewMat(2, 3, Elem{2, 1, 2, 4, 5, 6})
// 	b := NewMat(2, 3, Elem{3, 6, 4, 2, 6, 7})

// 	a.Add(b)

// 	println(a.String())
// }
