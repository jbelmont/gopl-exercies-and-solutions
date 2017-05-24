package shaHashDiff

import (
	"testing"
)

func TestGenerateShahash(t *testing.T) {
	c1 := generateShahash("s")

	if c1 == "" {
		t.Errorf("%s should not be nil", c1)
	}
}

func TestCompareBits(t *testing.T) {
	c1 := generateShahash("s")
	c2 := generateShahash("S")

	diff := compareBits(c1, c2)
	if diff == 0 {
		t.Error("should not be equal zero")
	}
}
