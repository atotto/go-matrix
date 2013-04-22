package matrix_test

import (
	. "."
	"testing"
)

func TestM(t *testing.T) {
	var mat Matrix
	mat = NewFloatMatrixFromElements(3, 2, []float64{1.2, 2.3, 4.5, 1.3, 4.4, 6.1})
	mat.Size()
}

func TestAt(t *testing.T) {
	mat := NewFloatMatrixFromElements(3, 2, []float64{1.2, 2.3, 4.5, 1.3, 4.4, 6.1})
	out := mat.At(1, 1)
	if out != 1.3 {
		t.Errorf("mat.At(1.1) want 1.3, got %v", out)
	}
}

func TestSet(t *testing.T) {
	mat := NewFloatMatrixFromElements(3, 2, []float64{1.2, 2.3, 4.5, 1.3, 4.4, 6.1})
	mat.Set(1, 1, 11.4)
	out := mat.At(1, 1)
	if out != 11.4 {
		t.Errorf("mat.At(1.1) want 11.4, got %v", out)
	}
}

func TestNew(t *testing.T) {
	mat := NewFloatMatrix(12, 34)
	elem := mat.Elem()

	if len(elem) != 12*34 {
		t.Errorf("length of the elements want %v, got %v", 12*34, len(elem))
	}
}
