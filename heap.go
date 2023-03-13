package algo

func MinHeapify(arr []int, n int, i int) {
	smallest := i
	l := 2*i + 1
	r := 2*i + 2

	if l < n && arr[l] < arr[smallest] {
		smallest = l
	}
	if r < n && arr[r] < arr[smallest] {
		smallest = r
	}
	if smallest != i {
		arr[smallest], arr[i] = arr[i], arr[smallest]
	}
}

func BinaryToMinHeap(arr []int) {
	for x := len(arr)/2 - 1; x >= 0; x-- {
		MinHeapify(arr, len(arr), x)
	}
}
