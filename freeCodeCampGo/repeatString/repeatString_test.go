package repeatString

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test(t *testing.T) {
  assert.Equal(t, Repeat(words: "*", num: 3).repeat, "***", "1")
	assert.Equal(t, Repeat(words: "abc", num: 3).repeat, "abcabcabc", "2")
	assert.Equal(t, Repeat(words: "abc", num: 4).repeat, "abcabcabcabc", "3")
	assert.Equal(t, Repeat(words: "abc", num: 1).repeat, "abc", "4")
	assert.Equal(t, Repeat(words: "*", num: 8).repeat, "********", "5")
	assert.Equal(t, Repeat(words: "abc", num: -2).repeat, "", "6")
}
