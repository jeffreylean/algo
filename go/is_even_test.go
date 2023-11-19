package algo

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsEven(t *testing.T) {
	assert.Equal(t, true, IsEven(2))
	assert.Equal(t, false, IsEven(3))
}
