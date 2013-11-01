// Addition
package WikiArith

func Add(numbers []float64) float64 {
	sum := 0.0
	for _, v := range numbers {
		sum += v
	}
	return sum
}
