package algo

func NextGreater(arr1 []int, arr2 []int) []int {
	stack := make([]int, 0)
	mp := make(map[int]int)
	res := make([]int, 0)

	for _, each := range arr2 {
		for len(stack) > 0 && stack[len(stack)-1] < each {
			mp[stack[len(stack)-1]] = each
			stack = stack[:len(stack)-1]
		}
		stack = append(stack, each)
	}

	for _, each := range arr1 {
		if v, ok := mp[each]; !ok {
			res = append(res, -1)
		} else {
			res = append(res, v)
		}
	}

	return res

}
