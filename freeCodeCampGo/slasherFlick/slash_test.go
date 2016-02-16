package slash

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test(t *testing.T) {
	assert.Equal(t, slasher([]int{1, 2, 3}, 2), []int{3}, "1")
	assert.Equal(t, slasher([]int{1, 2, 3}, 0), []int{1, 2, 3}, "2")
	assert.Equal(t, slasher([]int{1, 2, 3}, 9), []int{}, "3")
	assert.Equal(t, slasher([]int{1, 2, 3}, 4), []int{}, "4")
	assert.Equal(t, slasher([]int{1, 2, 3, 4, 5}, 3), []int{4, 5}, "5")
}
