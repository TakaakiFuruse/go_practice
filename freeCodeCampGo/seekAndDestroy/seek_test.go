package seek

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test(t *testing.T) {
	assert.Equal(t, destroyer([]int{1, 2, 3, 1, 2, 3}, 2, 3), []int{1, 1}, "1")
	assert.Equal(t, destroyer([]int{1, 2, 3, 5, 1, 2, 3}, 2, 3), []int{1, 5, 1}, "2")
	assert.Equal(t, destroyer([]int{3, 5, 1, 2, 2}, 2, 3, 5), []int{1}, "")
	assert.Equal(t, destroyer([]int{2, 3, 2, 3}, 2, 3), []int{}, "")
	// assert.Equal(t, destroyer([]int{"tree", "hamburger", 53}, "tree", 53), []int{"hamburger"}, "")
}
