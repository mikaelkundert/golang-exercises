package tempconv

import (
	"fmt"
)

// Print returns given t value in string
func Print(t float64) {
	c := Celsius(t)
	fmt.Printf("%s is:\n", c)
	fmt.Printf("\t%s\n", CToF(c))
	fmt.Printf("\t%s\n", CToK(c))

	f := Fahrenheit(t)
	fmt.Printf("%s is:\n", f)
	fmt.Printf("\t%s\n", FToC(f))
	fmt.Printf("\t%s\n", FToK(f))

	k := Kelvin(t)
	fmt.Printf("%s is:\n", k)
	fmt.Printf("\t%s\n", KToC(k))
	fmt.Printf("\t%s\n", KToF(k))
}
