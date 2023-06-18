package algo

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLRUPutAndGet(t *testing.T) {
	lru := NewLRU(5)

	lru.Put(1, 4)
	lru.Put(2, 5)

	assert.Equal(t, 4, lru.Get(1))
	assert.Equal(t, 5, lru.Get(2))
}

func TestLRUOverwriting(t *testing.T) {
	lru := NewLRU(5)

	lru.Put(1, 4)
	assert.Equal(t, 4, lru.Get(1))
	lru.Put(1, 5)
	assert.Equal(t, 5, lru.Get(1))
}

func TestLRUOverSize(t *testing.T) {
	lru := NewLRU(5)

	lru.Put(1, 4)
	lru.Put(2, 2)
	lru.Put(3, 6)
	lru.Put(4, 1)
	lru.Put(5, 3)

	assert.Equal(t, 4, lru.Get(1))
	// Key 1 should be the least used, if new cache comes in, 1 should be removed
	lru.Put(6, 2)
	assert.Equal(t, -1, lru.Get(1))
	assert.Equal(t, 2, lru.Get(6))
}
