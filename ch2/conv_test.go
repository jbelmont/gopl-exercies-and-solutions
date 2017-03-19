package tempconv

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCToF(t *testing.T) {
	actual := float64(CToF(45))
	expected := float64(113)
	assert.Equal(t, actual, expected)
}

func TestFToC(t *testing.T) {
	actual := float64(FToC(45))
	expected := float64(7.222222222222222)
	assert.Equal(t, actual, expected)
}

func TestCToK(t *testing.T) {
	actual := float64(CToK(45))
	expected := float64(318150)
	assert.Equal(t, actual, expected)
}

func TestKToC(t *testing.T) {
	actual := float64(KToC(45))
	expected := float64(273.195)
	assert.Equal(t, actual, expected)
}
