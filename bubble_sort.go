package algo

// Most slowest sorting algo out there.
// Time complexity of O(N^2)
// Because of the 2 layer nested loop.
func BubbleSort(arr []int) []int {
	for x := 0; x < len(arr)-1; x++ {
		for j := 0; j < len(arr)-x-1; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}
	return arr
}
