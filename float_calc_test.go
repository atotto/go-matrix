package matrix

import (
	"testing"
)

func TestMul(t *testing.T) {

	type mulTest struct {
		a      *FloatMatrix
		b      *FloatMatrix
		expect *FloatMatrix
	}

	var mulTests = []mulTest{
		mulTest{
			NewFloatMatrixFromElements(1, 3, []float64{1, 2, 3}),
			NewFloatMatrixFromElements(3, 1, []float64{1, 2, 3}),
			NewFloatMatrixFromElements(1, 1, []float64{14}),
		},
		mulTest{
			NewFloatMatrixFromElements(3, 1, []float64{1, 2, 3}),
			NewFloatMatrixFromElements(1, 3, []float64{1, 2, 3}),
			NewFloatMatrixFromElements(3, 3, []float64{1, 2, 3, 2, 4, 6, 3, 6, 9}),
		},
	}

	for _, test := range mulTests {
		out := test.a.Mul(test.b)
		if !out.Eq(test.expect) {
			t.Errorf("Mul want:\n%s\na:\n%s\nb:\n%s\ngot:\n%s", test.expect, test.a, test.b, out)
		}
	}
}
