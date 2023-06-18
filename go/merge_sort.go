package algo

func MergeSort(array []int) []int {
	if len(array) < 2 {
		return array
	}

	first := MergeSort(array[:len(array)/2])
	second := MergeSort(array[len(array)/2:])
	return MergeArray(first, second)
}

func MergeArray(a []int, b []int) []int {
	sorted := make([]int, 0)
	i := 0
	j := 0

	for i < len(a) && j < len(b) {
		if a[i] < b[j] {
			sorted = append(sorted, a[i])
			i++
		} else {
			sorted = append(sorted, b[j])
			j++
		}
	}

	for ; i < len(a); i++ {
		sorted = append(sorted, a[i])
	}
	for ; j < len(b); j++ {
		sorted = append(sorted, b[j])
	}
	return sorted
}
