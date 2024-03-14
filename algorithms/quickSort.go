package algorithms

import (
	"fmt"
)

func swap(arr []int, x int, y int) {
	arr[x], arr[y] = arr[y], arr[x]
}

func partition(arr []int, low int, high int) int {
	pivot := arr[high]
	index := low - 1

	for iterator := low; iterator < high; iterator++ {
		if arr[iterator] < pivot {
			index++
			swap(arr, index, iterator)
		}
	}

	index++
	swap(arr, index, high)

	return index
}

func sort(arr []int, low int, high int) {
	if low < high {
		pivot := partition(arr, low, high)

		sort(arr, low, pivot-1)
		sort(arr, pivot+1, high)
	}

}

func QuickSort(arr []int) {
	sort(arr, 0, len(arr)-1)
	fmt.Println(arr)
}
