package algo

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMinimumSwaps(t *testing.T) {
	arr := []int32{1, 4, 3, 2}

	swaps := MinimumSwaps(arr)
	assert.Equal(t, int32(1), swaps)
}
