package algorithms

import "fmt"

func InsertionSort(array []int) {
	arrayLength := len(array)
	for iterator := 1; iterator < arrayLength; iterator++ {
		currentValue := array[iterator]
		index := iterator - 1
		for index >= 0 && array[index] > currentValue {
			array[index+1] = array[index]
			index--
		}
		array[index+1] = currentValue
	}
	fmt.Println(array)
}
