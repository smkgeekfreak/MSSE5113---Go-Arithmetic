// Subtraction_test.go
package WikiArith

import (
	"testing"
)

func TestSubtract(t *testing.T) {
	in := []float64{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	out := -55.0
	if x := Subtract(in); x != out {
		t.Errorf("Add(%v) = %v, wanted %v", in, x, out)
	}
}

func TestSubtractNeg(t *testing.T) {
	in := []float64{1, -2, 3, -4, 5, 6, 7, 0, 9, 0}
	out := -25.0
	if x := Subtract(in); x != out {
		t.Errorf("Add(%v) = %v, wanted %v", in, x, out)

	}
}
func TestSubtractZeros(t *testing.T) {
	in := []float64{0, 0, 0, 0, 0}
	out := 0.0
	if x := Subtract(in); x != out {
		t.Errorf("Add(%v) = %v, wanted %v", in, x, out)

	}
}
