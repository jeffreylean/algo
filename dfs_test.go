package algo

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMinimumSwaps(t *testing.T) {
	arr := []int32{4, 3, 1, 2}

	swaps := MinimumSwaps(arr)
	assert.Equal(t, int32(3), swaps)
}
