package algorithms

import "fmt"

func FibonacciForLoop(limit int) {
	secondPreviousNumber, firstPreviousNumber := 0, 1

	fmt.Println(secondPreviousNumber)
	fmt.Println(firstPreviousNumber)

	for counter := 0; counter < limit-2; counter++ {
		temporary_value := secondPreviousNumber + firstPreviousNumber

		fmt.Println(temporary_value)

		secondPreviousNumber = firstPreviousNumber
		firstPreviousNumber = temporary_value
	}
}

func FibonacciRecursion(limit int) {
	counter, secondPreviousNumber, firstPreviousNumber := 0, 0, 1

	fmt.Println(secondPreviousNumber)
	fmt.Println(firstPreviousNumber)

	var fibonacci func(int, int)
	fibonacci = func(secondPreviousNumber int, firstPreviousNumber int) {

		temporary_value := secondPreviousNumber + firstPreviousNumber

		fmt.Println(temporary_value)

		secondPreviousNumber = firstPreviousNumber
		firstPreviousNumber = temporary_value

		counter++
		if counter < limit-2 {
			fibonacci(secondPreviousNumber, firstPreviousNumber)
		}
	}

	fibonacci(secondPreviousNumber, firstPreviousNumber)
}

func FibonacciNth(nthNumber int) {
	var nth func(int) int
	nth = func(number int) int {
		if number <= 1 {
			return number
		}
		return nth(number-2) + nth(number-1)
	}

	fmt.Println(nth(nthNumber - 1))
}
