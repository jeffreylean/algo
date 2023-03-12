package algo

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStringToByte(t *testing.T) {
	s := "hello"
	assert.Equal(t, []byte(s), StringToByte(s))
}
