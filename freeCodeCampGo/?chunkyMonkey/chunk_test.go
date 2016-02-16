package chunk

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test(t *testing.T) {

	// a1 := [][]string{[]string{"a", "b"}, []string{"c", "d"}}
	a2 := [][]int{[]int{0, 1, 2}, []int{3, 4, 5}}
	a3 := [][]int{[]int{0, 1}, []int{2, 3}, []int{4, 5}}
	a4 := [][]int{[]int{0, 1, 2, 3}, []int{4, 5}}
	a5 := [][]int{[]int{0, 1, 2}, []int{3, 4, 5}, []int{6}}
	a6 := [][]int{[]int{0, 1, 2, 3}, []int{4, 5, 6, 7}, []int{8}}

	// assert.Equal(t, chunk([]string{"a", "b", "c", "d"}, 2), a1, "1")
	assert.Equal(t, chunk([]int{0, 1, 2, 3, 4, 5}, 3), a2, "2")
	assert.Equal(t, chunk([]int{0, 1, 2, 3, 4, 5}, 2), a3, "3")
	assert.Equal(t, chunk([]int{0, 1, 2, 3, 4, 5}, 4), a4, "4")
	assert.Equal(t, chunk([]int{0, 1, 2, 3, 4, 5, 6}, 3), a5, "5")
	assert.Equal(t, chunk([]int{0, 1, 2, 3, 4, 5, 6, 7, 8}, 4), a6, "6")

}
