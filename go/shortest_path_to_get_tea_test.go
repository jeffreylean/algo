package algo

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetBubbleTea(t *testing.T) {
	grid := [][]byte{[]byte("XXXXXX"), []byte("X*OOOX"), []byte("XOO#OX"), []byte("XXXXXX")}

	distance := GetBubbleTea(grid)

	assert.Equal(t, 3, distance)
}
