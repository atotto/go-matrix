package matrix_test

import (
	. "."
	"image"
	"testing"
)

func TestSize(t *testing.T) {
	var mat Matrix
	mat = NewFloatMatrixFromElements(3, 2, []float64{1.2, 2.3, 4.5, 1.3, 4.4, 6.1})
	m, n := mat.Size()
	if m != 3 || n != 2 {
		t.Errorf("Size() want 3 2 , got %d %d", m, n)
	}
}

func TestRow(t *testing.T) {
	var mat Matrix
	mat = NewFloatMatrixFromElements(3, 2, []float64{1.2, 2.3, 4.5, 1.3, 4.4, 6.1})
	m := mat.Row()

	if m != 3 {
		t.Errorf("Row() want 3 , got %d", m)
	}
}

func TestCol(t *testing.T) {
	var mat Matrix
	mat = NewFloatMatrixFromElements(3, 2, []float64{1.2, 2.3, 4.5, 1.3, 4.4, 6.1})
	n := mat.Col()

	if n != 2 {
		t.Errorf("Row() want 2 , got %d", n)
	}
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

func TestFind(t *testing.T) {
	mat := NewFloatMatrixFromElements(3, 3, []float64{1, 2, 3, 4, 5, 6, 7, 8, 9})

	p := make(chan image.Point)
	var pt image.Point

	go func() {
		for {
			pt = <-p
		}
	}()

	mat.Find(8, p)

	if pt.X != 2 || pt.Y != 1 {
		t.Errorf("element i,j want 2,1, got %d,%d", pt.X, pt.Y)
	}
}

func TestEq(t *testing.T) {
	type eqTest struct {
		a      *FloatMatrix
		b      *FloatMatrix
		expect bool
	}

	var eqTests = []eqTest{
		eqTest{
			NewFloatMatrixFromElements(3, 2, []float64{1, 2, 3, 4, 5, 6}),
			NewFloatMatrixFromElements(3, 2, []float64{1, 2, 3, 4, 5, 6}),
			true},
		eqTest{
			NewFloatMatrixFromElements(3, 2, []float64{1, 2, 3, 4, 5, 6}),
			NewFloatMatrixFromElements(3, 2, []float64{1, 3, 3, 4, 5, 6}),
			false},
		eqTest{
			NewFloatMatrixFromElements(3, 2, []float64{1, 2, 3, 4, 5, 6}),
			NewFloatMatrixFromElements(2, 3, []float64{1, 2, 3, 4, 5, 6}),
			false},
		eqTest{
			NewFloatMatrixFromElements(1, 2, []float64{1, 2}),
			NewFloatMatrixFromElements(1, 3, []float64{1, 2, 3}),
			false},
	}

	for _, test := range eqTests {
		out := test.a.Eq(test.b)
		if out != test.expect {
			t.Errorf("Eq want %t, got %t\na mat is:\n%s\nb mat is:\n%s", test.expect, out, test.a.String(), test.b.String())
		}
	}
}

func TestTranspose(t *testing.T) {
	type transposeTest struct {
		mat    *FloatMatrix
		expect *FloatMatrix
	}

	var transposeTests = []transposeTest{
		transposeTest{
			NewFloatMatrixFromElements(3, 2, []float64{1, 2, 3, 4, 5, 6}),
			NewFloatMatrixFromElements(2, 3, []float64{1, 3, 5, 2, 4, 6}),
		},
	}

	for _, test := range transposeTests {
		out := test.mat.Transpose()
		if !out.Eq(test.expect) {
			t.Errorf("Transpose want:\n%s\ngot:\n%s", test.expect.String(), out.String())
		}
	}
}
