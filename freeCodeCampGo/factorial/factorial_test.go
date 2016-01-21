package factorial

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFactorialize(t *testing.T) {
	assert.Equal(t, factorialize(5), 12)
	assert.Equal(t, factorialize(10), 3628800)
	assert.Equal(t, factorialize(20), 2432902008176640000)
	assert.Equal(t, factorialize(0), 1)
}
