package algorithms

import "fmt"

func BubbleSort(array []int) {
	arrayLength := len(array)
	fmt.Println(array)
	fmt.Println("---------")
	for iterator := 0; iterator < arrayLength-1; iterator++ {
		fmt.Printf("[%v] -> %v ", iterator, array)
		swapped := false
		for index := 0; index < arrayLength-iterator-1; index++ {
			if array[index] > array[index+1] {
				fmt.Printf("(%v-%v),", array[index], array[index+1])
				array[index+1], array[index] = array[index], array[index+1]
				swapped = true
			}
		}
		fmt.Print("\n")
		if !swapped {
			break
		}
	}
	fmt.Println("---------")
	fmt.Println(array)
}
