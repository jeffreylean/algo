package algo

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBinarySearch(t *testing.T) {
	arr := []int{0, 3, 6, 8, 9}

	i := BinarySearch(arr, 3)
	assert.Equal(t, 1, i)
	assert.NotEqual(t, 2, i)
}
