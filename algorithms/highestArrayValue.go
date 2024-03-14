package algorithms

import "fmt"

func HighestArrayValue(array []int) {
	highestValue := array[0]

	for index := 0; index < len(array); index++ {
		if array[index] > highestValue {
			highestValue = array[index]
		}
	}
	fmt.Println(highestValue)
}
