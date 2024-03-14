package algorithms

import "fmt"

func LowestArrayValue(array []int) {
	lowestValue := array[0]

	for index := 0; index < len(array); index++ {
		if array[index] < lowestValue {
			lowestValue = array[index]
		}
	}
	fmt.Println(lowestValue)
}
