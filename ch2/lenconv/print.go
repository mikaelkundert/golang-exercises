package lenconv

import "fmt"

// Print outputs given value in variety of lengths
func Print(t float64) {
	f := Feet(t)
	fmt.Printf("%s is:\n", f)
	fmt.Printf("\t%s\n", FToM(f))

	m := Meter(t)
	fmt.Printf("%s is:\n", m)
	fmt.Printf("\t%s\n", MToF(m))
}
