// Package tempconv performs Celsius and Fahrenheit conversions.
package tempconv

import "fmt"

// Celsius is a float64 value
type Celsius float64

// Fahrenheit is a float64 value
type Fahrenheit float64

// Kelvin is a float64 value
type Kelvin float64

const (
	AbsoluteZeroC Celsius = -273.15
	FreezingC     Celsius = 0
	BoilingC      Celsius = 100
	ZeroK         Kelvin  = -273.15
)

func (c Celsius) String() string {
	return fmt.Sprintf("%g°C", c)
}

func (f Fahrenheit) String() string {
	return fmt.Sprintf("%g°F", f)
}

func (k Kelvin) String() string {
	return fmt.Sprintf("%g°K", k)
}

// Exercise 2.1: Add types, constants, and functions to tempconv for processing temperatures in the Kelvin scale,
// where zero Kelvin is −273.15°C and a difference of 1K has the same magnitude as 1°C.
