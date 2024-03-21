package algo

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNextGreater(t *testing.T) {
	num1 := []int{5}
	num2 := []int{5, 4, 3, 2, 9}

	assert.Equal(t, NextGreater(num1, num2), []int{9})
}
