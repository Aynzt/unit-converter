// Package temperature perfom Celsuis and Fahrenheint conversions
package temperature

import "fmt"

type Celsuis float64
type Fahrenheit float64
type Kelvin float64

const (
	AbsoluteZeroC Celsuis = -273.15
	FreezingC     Celsuis = 0
	BoilingC      Celsuis = 100
	AbsoluteZeroK Kelvin  = Kelvin(AbsoluteZeroC)
)

func (c Celsuis) String() string {
	return fmt.Sprintf("%g°C", c)
}

func (f Fahrenheit) String() string {
	return fmt.Sprintf("%g°F", f)
}

func (k Kelvin) String() string {
	return fmt.Sprintf("%g°K", k)
}
