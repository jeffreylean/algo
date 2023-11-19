package algo

func add() int {
	n := 6
	sum := 0
	iter := 1

	for i := 1; i <= n; i++ {
		if i == iter*iter {
			iter++
			continue
		}
		sum += i
	}
	return sum
}
