package arrays // Refactor - Refatoração / for range / _ blank identifier (identificador em branco)

// Sum retona o cálculo da somas dos elementos de uma matriz de 5 números inteiros.
func Sum(numbers [5]int) int {
	var sum int = 0
	for _, number := range numbers {
		sum += number
	}
	return sum
}

/*	Iteração / Laços de Repetições utilizando range
	https://go.dev/doc/effective_go#for
	_ blank identifier (identificador em branco)
	https://go.dev/doc/effective_go#blank
*/
