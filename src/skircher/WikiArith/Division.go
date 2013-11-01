// Division
package WikiArith

func Divide(numbers []float64) float64 {
	var result float64
	for i, v := range numbers {
		if i == 0 {
			result = v
		} else {
			result = result / v
		}
	}
	return result
}
