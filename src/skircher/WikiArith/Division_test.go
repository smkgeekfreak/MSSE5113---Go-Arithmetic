// Division_test
package WikiArith

import (
	"math"
	"testing"
)

func TestDivide(t *testing.T) {
	in := []float64{10.0, 2.0, 3.0}
	out := 1.666667
	if x := Divide(in); math.Abs(out-x) >= 0.000001 {
		t.Errorf("Divide(%f) = %f, wanted %f", in, x, out)
	}
}

func TestDivideNeg(t *testing.T) {
	in := []float64{10.0, -2.532}
	out := -3.949447
	if x := Divide(in); math.Abs(x-out) > 0.000001 {
		t.Errorf("Divide(%f) = %f, wanted %f", in, x, out)
	}
}

//func TestDivideEvenNegs(t *testing.T) {
//	in := []int{1, -2, 3, -4, 5, 6, 7, 4, 9, 6}
//	out := 1088640.0
//	if x := Divide(in); x != out {
//		t.Errorf("Divide(%v) = %v, wanted %v", in, x, out)

//	}
//}

//func TestDivideZeros(t *testing.T) {
//	in := []int{0, 1, 2, 3, 4}
//	out := 0
//	if x := Divide(in); x != out {
//		t.Errorf("Divide(%v) = %v, wanted %v", in, x, out)

//	}
//}
