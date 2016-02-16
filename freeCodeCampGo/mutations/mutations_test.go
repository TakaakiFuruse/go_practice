package mutations

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test(t *testing.T) {
	assert.Equal(t, mutation([]string{"hello", "hey"}), false, "1")
	assert.Equal(t, mutation([]string{"hello", "Hello"}), true, "2")
	assert.Equal(t, mutation([]string{"zyxwvutsrqponmlkjihgfedcba", "qrstu"}), true, "3")
	assert.Equal(t, mutation([]string{"Mary", "Army"}), true, "4")
	assert.Equal(t, mutation([]string{"Mary", "Aarmy"}), true, "5")
	assert.Equal(t, mutation([]string{"Alien", "line"}), true, "6")
	assert.Equal(t, mutation([]string{"floor", "for"}), true, "7")
	assert.Equal(t, mutation([]string{"hello", "neo"}), false, "8")
}
