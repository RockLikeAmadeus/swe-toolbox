package main

func Sum(numbers []int) int {
	sum := 0
	for _, number := range numbers {
		sum += number
	}
	return sum
}

// Variadic functions take a variable number of arguments
func SumAll(numbersToSum ...[]int) []int {
	/*
		lengthOfNumbers := len(numbersToSum)

		// Create a slice with a starting capacity with `make()``
		sums := make([]int, lengthOfNumbers)

		for i, numbers := range numbersToSum {
			sums[i] = Sum(numbers)
		}
	*/

	var sums []int
	for _, numbers := range numbersToSum {
		// `append()` returns a new slice
		sums = append(sums, Sum(numbers))
	}

	return sums
}

func SumAllTails(numbersToSum ...[]int) []int {
	var sums []int
	for _, numbers := range numbersToSum {
		if len(numbers) == 0 {
			sums = append(sums, 0)
		} else {
			tail := numbers[1:]
			sums = append(sums, Sum(tail))
		}
	}

	return sums
}
