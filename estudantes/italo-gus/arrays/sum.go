package arrays // Arrays

// Sum retona o cálculo da somas dos elementos de uma matriz de 5 números inteiros.
func Sum(numbers [5]int) int {
	var sum int = 0
	for i := 0; i < 5; i++ {
		sum += numbers[i]
	}
	return sum
}
