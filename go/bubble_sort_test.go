package algo

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBubbleSort(t *testing.T) {
	arr := []int{3, 2, 5, 8, 1}

	assert.ElementsMatch(t, []int{1, 2, 3, 5, 8}, BubbleSort(arr))
}
