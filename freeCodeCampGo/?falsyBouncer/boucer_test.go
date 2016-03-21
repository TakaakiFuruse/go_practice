package bouncer

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test(t *testing.T) {
	assert.Equal(t, bouncer(any{7, "ate", "", false, 9}), any{7, "ate", 9}, "1")
	assert.Equal(t, bouncer(any{"a", "b", "c"}), any{"a", "b", "c"}, "2")
	assert.Equal(t, bouncer(any{false, nil, 0, nil, nil, ""}), any{}, "3")
}
