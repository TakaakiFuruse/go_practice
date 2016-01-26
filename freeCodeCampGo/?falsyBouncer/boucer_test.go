import (
  "testing"

  "github.com/stretchr/testify/assert"
)

func Test(t *testing.T) {
assert.Equal(t, bouncer([7, "ate", "", false, 9]), [7, "ate", 9], "1")
assert.Equal(t, bouncer(["a", "b", "c"]), ["a", "b", "c"], "2")
assert.Equal(t, bouncer([false, null, 0, NaN, undefined, ""]), [], "3")

}

