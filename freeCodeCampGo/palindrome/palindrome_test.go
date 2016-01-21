package palindrome

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test(t *testing.T) {
	assert.Equal(t, palindrome("eye"), true, "1")
	assert.Equal(t, palindrome("race car"), true, "2")
	assert.Equal(t, palindrome("not a palindrome"), false, "3")
	assert.Equal(t, palindrome("A man, a plan, a canal. Panama"), true, "4")
	assert.Equal(t, palindrome("never odd or even"), true, "5")
	assert.Equal(t, palindrome("nope"), false, "6")
	assert.Equal(t, palindrome("almostomla"), false, "7")
	assert.Equal(t, palindrome("My age is 0, 0 si ega ym."), true, "8")
	assert.Equal(t, palindrome("1 eye for of 1 eye."), false, "9")
}
