package confEnd

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test(t *testing.T) {
	assert.Equal(t, end("Bastian", "n"), true, "1")
	// assert.Equal(t, end("Connor", "n"), false, "2")
	// assert.Equal(t, end("Walking on water and developing software from a specification are easy if both are frozen", "specification"), false, "3")
	// assert.Equal(t, end("He has to give me a new name", "name"), true, "4")
	// assert.Equal(t, end("He has to give me a new name", "me"), true, "5")
	// assert.Equal(t, end("He has to give me a new name", "na"), false, "6")
	// assert.Equal(t, end("If you want to save our world, you must hurry. We dont know how much longer we can withstand the nothing", "mountain"), false, "7")
}
