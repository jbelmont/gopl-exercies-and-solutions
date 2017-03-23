package main

import (
	"math"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestF(t *testing.T) {
	const PI = math.Pi
	val := f(PI, 1.5)
	assert.Equal(t, val, -0.09571998986588899)
}

func TestCorner(t *testing.T) {
	i, j := 2, 3
	sx, sy := corner(i, j)
	assert.Equal(t, sx, 297.4019237886467)
	assert.Equal(t, sy, 11.373325050200139)
}
