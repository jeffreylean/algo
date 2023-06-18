package algo

// Find minimum number of swaps to sort an array.
// Solution: use depth first search to find the minimum swaps.

func MinimumSwaps(arr []int32) int32 {
	m := make(map[int]int, len(arr))
	var swapCounts int32
	// Load into maps with indexes as key, and it's value of the indexes as value
	for i, each := range arr {
		m[i] = int(each)
	}
	// List of boolean to record which node has been visited
	visited := make([]bool, len(arr))

	for i := range arr {
		// Mark current point to visited
		visited[i] = true

		// Check if current point is sorted by checking if the value of indexes is indexes+1
		// arr[i] = i+1 == sorted
		for m[i] != i+1 {
			// Not sorted, than swap it to the index where it should belong.
			target := m[i] - 1
			temp := m[target]
			m[i], m[target] = temp, m[i]

			// Mark this newly swapped index to be visited
			visited[target] = true
			swapCounts++
		}
	}

	return swapCounts
}
