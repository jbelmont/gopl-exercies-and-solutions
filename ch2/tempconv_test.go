package tempconv

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestString(t *testing.T) {
	celsius := Celsius(1.75).String()
	celsiusExpected := "1.75°C"
	assert.Equal(t, celsius, celsiusExpected)

	fahrenheit := Fahrenheit(1.75).String()
	fahrenheitExpected := "1.75°F"
	assert.Equal(t, fahrenheit, fahrenheitExpected)

	kelvin := Kelvin(5.5).String()
	kelvinExpected := "5.5°K"
	assert.Equal(t, kelvin, kelvinExpected)
}
