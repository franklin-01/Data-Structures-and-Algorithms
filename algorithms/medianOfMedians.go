package algorithms

func findMedian(array []int, right int, left int) int {
	size := right - left + 1
	if size%2 == 0 {
		return array[left+size/2-1]
	}
	return array[right+size/2-1]
}

func Swap(array []int, first_position int, second_position int) {
	array[first_position], array[second_position] = array[second_position], array[first_position]
}

func Partition(array []int, left int, right int, pivotIndex int) int {
	pivotValue := array[pivotIndex]
	Swap(array, pivotIndex, right)

	storeIndex := left

	for index := left; index < right; index++ {
		if array[index] < pivotValue {
			Swap(array, storeIndex, index)
			storeIndex++
		}
	}
	Swap(array, right, storeIndex)

	return storeIndex
}

func MedianOfMedians(array []int, left int, right int) int {
	if left == right {
		return array[left]
	}

	for {
		if left <= right {
			size := right - left + 1
			groups := (size + 4) / 5
			medians := make([]int, groups)

			for iterator := 0; iterator < groups; iterator++ {
				groupLeft := left + iterator*5
				groupRight := groupLeft + 4
				if groupRight > right {
					groupRight = right
				}

				medians[iterator] = findMedian(array, groupLeft, groupRight)
			}

			pivotIndex := MedianOfMedians(medians, 0, groups-1)
			pivotIndex = Partition(array, left, right, pivotIndex)

			return pivotIndex
		}
	}
}
