package algo

// QuickSortDivideAndConquer: Quick sorting with divide and conquer method, this algo does
// not handle duplication
func QuickSortDivideAndConquer(arr []int) []int {
	// base case
	if len(arr) < 2 {
		return arr
	} else {
		// Recursive case
		smaller := make([]int, 0)
		larger := make([]int, 0)

		//pivot
		pivot := arr[0]
		for _, value := range arr {
			if value < pivot {
				smaller = append(smaller, value)
			}
			if value > pivot {
				larger = append(larger, value)
			}
		}
		ans := make([]int, 0)
		ans = append(ans, QuickSortDivideAndConquer(smaller)...)
		ans = append(ans, pivot)
		ans = append(ans, QuickSortDivideAndConquer(larger)...)
		return ans
	}
}

// QuickSortPartition: Quick sort that handle duplication by using Haore partition algo
func QuickSortPartition(arr []int, low int, high int) []int {
	if low < high {
		q := HoarePartition(&arr, low, high)
		QuickSortPartition(arr, low, q)
		QuickSortPartition(arr, q+1, high)
	}
	return arr
}

// HoarePartition: Haore partition algorithm
func HoarePartition(arr *[]int, low int, high int) int {
	pivot := (*arr)[low]
	i := low - 1
	j := high + 1
	for {
		for {
			j--
			if (*arr)[j] <= pivot {
				break
			}
		}

		for {
			i++
			if (*arr)[i] >= pivot {
				break
			}
		}

		if i < j {
			(*arr)[i], (*arr)[j] = (*arr)[j], (*arr)[i]
		} else {
			return j
		}
	}
}
