package tempconv

import "fmt"

// Celsius is a float64 value
type Celsius float64

// Fahrenheit is a float64 value
type Fahrenheit float64

// Kelvin is a float64 value
type Kelvin float64

const (
	// AbsoluteZeroC is constant
	AbsoluteZeroC Celsius = -273.15
	FreezingC     Celsius = 0
	BoilingC      Celsius = 100
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
