package caseSentence

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCaseSentence(t *testing.T) {
	assert.Equal(t, titleCase("I'm a little tea pot"), "I'm A Little Tea Pot", "1")
	assert.Equal(t, titleCase("sHoRt AnD sToUt"), "Short And Stout", "2")
	assert.Equal(t, titleCase("HERE IS MY HANDLE HERE IS MY SPOUT"), "Here Is My Handle Here Is My Spout", "3")
}
