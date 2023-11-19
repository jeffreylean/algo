package algo

func IsEven(n int) bool {
	if n&1 == 0 {
		return true
	}
	return false
}
