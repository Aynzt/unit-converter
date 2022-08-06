// Package weight perfom Kilogram and Pound conversions
package weight

import "fmt"

type Pound float64
type Kilogram float64

func (p Pound) String() string {
	return fmt.Sprintf("%g pound", p)
}

func (k Kilogram) String() string {
	return fmt.Sprintf("%g kg", k)
}
