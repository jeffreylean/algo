package algo

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFloodFill(t *testing.T) {
	img := [][]int{[]int{1, 1, 1}, []int{1, 1, 0}, []int{1, 0, 1}}
	assert.Equal(t, [][]int{[]int{2, 2, 2}, []int{2, 2, 0}, []int{2, 0, 1}}, floodFill(img, 1, 1, 2))
}
