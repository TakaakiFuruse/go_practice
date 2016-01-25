package largestNumber

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test(t *testing.T) {
	a1 := [][]int{
		[]int{13, 27, 18, 26},
		[]int{4, 5, 1, 3},
		[]int{32, 35, 37, 39},
		[]int{1000, 1001, 857, 1},
	}
	a2 := [][]int{
		[]int{4, 9, 1, 3},
		[]int{13, 35, 18, 26},
		[]int{32, 35, 97, 39},
		[]int{1000000, 1001, 857, 1},
	}
	assert.Equal(t, largestOfFour(a1), []int{27, 5, 39, 1001}, "1")
	assert.Equal(t, largestOfFour(a2), []int{9, 35, 97, 1000000}, "2")

}
