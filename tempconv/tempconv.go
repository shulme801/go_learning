// Package tempconv performs Celsius and Fahrenheit conversions.package tempconv
package tempconv

import "fmt"

// Celsius will be exported
type Celsius float64

// Fahrenheit will be exported
type Fahrenheit float64

// Exported Comments
const (
	AbsoluteZeroC Celsius = -273.15
	FreezingC     Celsius = 0
	BoilingC      Celsius = 100
)

func (c Celsius) String() string    { return fmt.Sprintf("%g°C", c) }
func (f Fahrenheit) String() string { return fmt.Sprintf("%g°F", f) }
