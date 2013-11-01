// Addition
package WikiArith

func Subtract(numbers []float64) float64 {
	diff := 0.0
	for _, v := range numbers {
		diff -= v
	}
	return diff
}
