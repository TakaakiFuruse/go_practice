package where

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test(t *testing.T) {
	assert.Equal(t, where([]int{10, 20, 30, 40, 50}, 35), 3, "")
	assert.Equal(t, where([]int{10, 20, 30, 40, 50}, 30), 2, "")
	assert.Equal(t, where([]int{40, 60}, 50), 1, "")
	assert.Equal(t, where([]int{3, 10, 5}, 3), 0, "")
	assert.Equal(t, where([]int{5, 3, 20, 3}, 5), 2, "")
	assert.Equal(t, where([]int{2, 20, 10}, 19), 2, "")
	assert.Equal(t, where([]int{2, 5, 10}, 15), 3., "")
}
