// Multiplication_test

package WikiArith

import (
	"testing"
)

func TestMultiply(t *testing.T) {
	in := []float64{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	out := 3628800.0
	if x := Multiply(in); x != out {
		t.Errorf("Multiply(%v) = %v, wanted %v", in, x, out)
	}
}

func TestMultiplyNeg(t *testing.T) {
	in := []float64{1, -2, 3, 4, 5, 6, 7, 4, 9, 6}
	out := -1088640.0
	if x := Multiply(in); x != out {
		t.Errorf("Multiply(%v) = %v, wanted %v", in, x, out)
	}
}

func TestMultiplyEvenNegs(t *testing.T) {
	in := []float64{1, -2, 3, -4, 5, 6, 7, 4, 9, 6}
	out := 1088640.0
	if x := Multiply(in); x != out {
		t.Errorf("Multiply(%v) = %v, wanted %v", in, x, out)

	}
}
func TestMultiplyZeros(t *testing.T) {
	in := []float64{0, 1, 2, 3, 4}
	out := 0.0
	if x := Multiply(in); x != out {
		t.Errorf("Multiply(%v) = %v, wanted %v", in, x, out)

	}
}
