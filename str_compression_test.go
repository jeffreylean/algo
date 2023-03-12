package algo

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCompression(t *testing.T) {
	i := Compress([]byte{'a', 'a', 'a', 'a'})
	assert.Equal(t, 2, i)
}
