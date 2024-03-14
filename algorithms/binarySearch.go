package algorithms

func BinarySearch(arr []int, value int) int {
	left, right := 0, len(arr)
	for left < right {
		mid := (left + right) / 2
		if arr[mid] == value {
			return arr[mid]
		}

		if arr[mid] < value {
			left = mid + 1
		} else {
			right = mid + 1
		}
	}

	return -1
}
