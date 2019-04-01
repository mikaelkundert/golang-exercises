// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

//!+

// Package tempconv performs Celsius and Fahrenheit conversions.
package tempconv

import "fmt"

// Celsius represents a temperature degree type.
type Celsius float64

// Fahrenheit represents a temperature degree type.
type Fahrenheit float64

// Kelvin represents a temperature degree type.
type Kelvin float64

const (
	// AbsoluteZeroC declares the absolute zero point for Celsius.
	AbsoluteZeroC Celsius = -273.15
	// FreezingC declares temperature when water gets freezing.
	FreezingC Celsius = 0
	// BoilingC declares temprature when water starts to boil.
	BoilingC Celsius = 100
)

func (c Celsius) String() string    { return fmt.Sprintf("%g°C", c) }
func (f Fahrenheit) String() string { return fmt.Sprintf("%g°F", f) }
func (k Kelvin) String() string     { return fmt.Sprintf("%g°F", k) }

//!-
