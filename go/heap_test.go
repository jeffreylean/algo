package algo

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMinHeapify(t *testing.T) {
	arr := []int{2, 3, 1, 4, 7, 9, 11}
	BinaryToMinHeap(arr)

	for x := 0; x < len(arr)/2; x++ {
		l := x*2 + 1
		r := x*2 + 2

		if l < len(arr) {
			assert.True(t, arr[x] <= l)
		}
		if r < len(arr) {
			assert.True(t, arr[x] <= r)
		}
	}
}
