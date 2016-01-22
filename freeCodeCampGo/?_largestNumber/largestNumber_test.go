package largestNumber

import (
  "testing"

  "github.com/stretchr/testify/assert"
)

func Test(t *testing.T) {
  a1 := []int{[]int{}}
  t1 := largestOfFour([[13, 27, 18, 26], [4, 5, 1, 3], [32, 35, 37, 39], [1000, 1001, 857, 1]])
  t2 := largestOfFour([[4, 9, 1, 3], [13, 35, 18, 26], [32, 35, 97, 39], [1000000, 1001, 857, 1]])
  assert.Equal(t, t1, [27,5,39,1001], "1")
assert.Equal(t, t2, [9, 35, 97, 1000000], "2")

}



