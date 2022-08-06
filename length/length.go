// Package length perfom Meter and Feet conversions
package length

import "fmt"

type Meter float64
type Feet float64

func (m Meter) String() string {
	return fmt.Sprintf("%g M", m)
}

func (f Feet) String() string {
	return fmt.Sprintf("%g FT", f)
}
