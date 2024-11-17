package main

func Sum(numbers []int) (sum int) {
	for _, number := range numbers {
		sum += number
	}
	return
}

// Go lets us write variadic functions,
// which means a func can take a variable number of arguments
func SumAll(numbersToSum ...[]int) (sums []int) {
	for _, numbers := range numbersToSum {
		sums = append(sums, Sum(numbers))
	}

	return
}

// Sums all items in a collection, except for the first one (the "head")
func SumAllTails(numbersToSum ...[]int) (sums []int) {
	for _, numbers := range numbersToSum {
		if len(numbers) == 0 { // empty slices have no tail!
			sums = append(sums, 0)
		} else {
			tail := numbers[1:] // slice[low:high]
			sums = append(sums, Sum(tail))
		}
	}
	return
}
