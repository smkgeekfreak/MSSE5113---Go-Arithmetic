// Multiplication
package WikiArith

func Multiply(numbers []float64) float64 {
	var result float64
	for i, v := range numbers {
		if i == 0 {
			result = v
		} else {
			result *= v
		}
	}
	return result
}
