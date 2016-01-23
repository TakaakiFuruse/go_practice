package truncate

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test(t *testing.T) {
	assert.Equal(t, truncate("A-tisket a-tasket A green and yellow basket", 11), "A-tisket...", "1")
	assert.Equal(t, truncate("Peter Piper picked a peck of pickled peppers", 14), "Peter Piper...", "2")
	assert.Equal(t, truncate("A-tisket a-tasket A green and yellow basket", 43), "A-tisket a-tasket A green and yellow basket", "3")
	assert.Equal(t, truncate("A-tisket a-tasket A green and yellow basket", 45), "A-tisket a-tasket A green and yellow basket", "4")
	assert.Equal(t, truncate("A-", 1), "A...", "5")
	assert.Equal(t, truncate("Absolutely Longer", 2), "Ab...", "6")
}
