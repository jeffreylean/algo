package algo

func BinarySearch(arr []int, target int) int {
	left := 0
	right := len(arr) - 1

	for left < right {
		mid := (left + right) / 2
		if target > arr[mid] {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}

	if left != len(arr) {
		if arr[left] == target {
			return left
		}
	}
	return -1
}
