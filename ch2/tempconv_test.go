package tempconv

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestString(t *testing.T) {
	celsius := Celsius(1.75)
	assert.Equal(t, reflect.TypeOf(celsius), reflect.TypeOf(Celsius(3)))

	fahrenheit := Fahrenheit(1.75)
	assert.Equal(t, reflect.TypeOf(fahrenheit), reflect.TypeOf(Fahrenheit(3.5)))

	kelvin := Kelvin(5.5)
	assert.Equal(t, reflect.TypeOf(kelvin), reflect.TypeOf(Kelvin(8)))
}
