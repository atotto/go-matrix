package matrix

import (
	"testing"
)

func TestMul(t *testing.T) {

	type mulTest struct {
		a      *Matrix[float64]
		b      *Matrix[float64]
		expect *Matrix[float64]
	}

	var mulTests = []mulTest{
		mulTest{
			NewMatrixFromElements[float64](1, 3, []float64{1, 2, 3}),
			NewMatrixFromElements[float64](3, 1, []float64{1, 2, 3}),
			NewMatrixFromElements[float64](1, 1, []float64{14}),
		},
		mulTest{
			NewMatrixFromElements[float64](3, 1, []float64{1, 2, 3}),
			NewMatrixFromElements[float64](1, 3, []float64{1, 2, 3}),
			NewMatrixFromElements[float64](3, 3, []float64{1, 2, 3, 2, 4, 6, 3, 6, 9}),
		},
	}

	for _, test := range mulTests {
		out := test.a.Mul(test.b)
		if !out.Eq(test.expect) {
			t.Errorf("Mul want:\n%s\na:\n%s\nb:\n%s\ngot:\n%s", test.expect, test.a, test.b, out)
		}
	}
}

func TestMulInt(t *testing.T) {

	type mulTest struct {
		a      *Matrix[int]
		b      *Matrix[int]
		expect *Matrix[int]
	}

	var mulTests = []mulTest{
		mulTest{
			NewMatrixFromElements[int](1, 3, []int{1, 2, 3}),
			NewMatrixFromElements[int](3, 1, []int{1, 2, 3}),
			NewMatrixFromElements[int](1, 1, []int{14}),
		},
		mulTest{
			NewMatrixFromElements[int](3, 1, []int{1, 2, 3}),
			NewMatrixFromElements[int](1, 3, []int{1, 2, 3}),
			NewMatrixFromElements[int](3, 3, []int{1, 2, 3, 2, 4, 6, 3, 6, 9}),
		},
	}

	for _, test := range mulTests {
		out := test.a.Mul(test.b)
		if !out.Eq(test.expect) {
			t.Errorf("Mul want:\n%s\na:\n%s\nb:\n%s\ngot:\n%s", test.expect, test.a, test.b, out)
		}
	}
}
