package algo

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestQuickSort(t *testing.T) {
	in := []int{123, 4545, 6, 4, 3, 23, 5, 6}

	out := QuickSortDivideAndConquer(in)
	assert.Equal(t, []int{3, 4, 5, 6, 23, 123, 4545}, out)

	out = QuickSortPartition(in, 0, len(in)-1)
	assert.Equal(t, []int{3, 4, 5, 6, 6, 23, 123, 4545}, out)
}
