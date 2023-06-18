package algo

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMergeSort(t *testing.T) {
	array := []int{4, 1, 34, 64, 5}
	assert.ElementsMatch(t, []int{1, 4, 5, 34, 64}, MergeSort(array))
}
