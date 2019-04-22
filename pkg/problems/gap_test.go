package problems

import (
	"math"
	"testing"
)
import "github.com/stretchr/testify/assert"

func TestFindMaxBinaryGap(t *testing.T) {
	assert.Equal(t, uint32(0), Solution(0))
	assert.Equal(t, uint32(0), Solution(1))
	assert.Equal(t, uint32(0), Solution(2))
	assert.Equal(t, uint32(0), Solution(3))
	assert.Equal(t, uint32(0), Solution(4))
	assert.Equal(t, uint32(1), Solution(5))
	assert.Equal(t, uint32(0), Solution(6))
	assert.Equal(t, uint32(0), Solution(7))
	assert.Equal(t, uint32(0), Solution(8))
	assert.Equal(t, uint32(2), Solution(9))

	assert.Equal(t, uint32(1), Solution(21))

	assert.Equal(t, uint32(0), Solution(32))

	assert.Equal(t, uint32(2), Solution(37))

	assert.Equal(t, uint32(4), Solution(529))

	assert.Equal(t, uint32(5), Solution(1041))

	assert.Equal(t, uint32(6), Solution(265126))

	assert.Equal(t, uint32(2), Solution(2147483547))

	assert.Equal(t, uint32(0), Solution(math.MaxInt32))
}