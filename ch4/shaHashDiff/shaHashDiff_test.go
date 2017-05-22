package shaHashDiff

import (
	"testing"
)

func TestGenerateShahash(t *testing.T) {
	c1 := generateShahash("s")
	c2 := generateShahash("S")

	if c1 == c2 {
		t.Errorf("%s should not equal %s", c1, c2)
	}
}
