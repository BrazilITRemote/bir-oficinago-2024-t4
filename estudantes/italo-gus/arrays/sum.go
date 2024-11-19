package arrays // Slice / Cobertura de teste no Go

// Sum retona o cálculo da somas dos elementos de um slice de números inteiros.
func Sum(numbers []int) int {
	var sum int = 0
	for _, number := range numbers {
		sum += number
	}
	return sum
}
