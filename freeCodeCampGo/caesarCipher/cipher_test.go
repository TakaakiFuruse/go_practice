package cipher

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test(t *testing.T) {
	assert.Equal(t, rot13("SERR PBQR PNZC"), "FREE CODE CAMP", "1")
	assert.Equal(t, rot13("SERR CVMMN!"), "FREE PIZZA!", "2")
	assert.Equal(t, rot13("SERR YBIR?"), "FREE LOVE?", "3")
	assert.Equal(t, rot13("GUR DHVPX OEBJA QBT WHZCRQ BIRE GUR YNML SBK."), "THE QUICK BROWN DOG JUMPED OVER THE LAZY FOX.", "4")
}
