package algorithms

import "fmt"

func SelectionSort(array []int) {
	length := len(array)
	fmt.Println(array)
	fmt.Println("---------")
	for iterator := 0; iterator < length-1; iterator++ {
		minIndex := iterator
		for index := iterator + 1; index < length; index++ {
			if array[index] < array[minIndex] {
				minIndex = index
			}
		}

		array[iterator], array[minIndex] = array[minIndex], array[iterator]
	}
	fmt.Println(array)
}
