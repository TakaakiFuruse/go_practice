package longestWord

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLongestWord(t *testing.T) {
	assert.Equal(t, findLongestWord("The quick brown fox jumped over the lazy dog"), 6, "1")
	assert.Equal(t, findLongestWord("May the force be with you"), 5, "2")
	assert.Equal(t, findLongestWord("Google do a barrel roll"), 6, "3")
	assert.Equal(t, findLongestWord("What is the average airspeed velocity of an unladen swallow"), 8, "4")
	assert.Equal(t, findLongestWord("What if we try a super-long word such as otorhinolaryngology"), 19)
}
