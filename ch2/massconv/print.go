package massconv

import "fmt"

// Print outputs the mass converstions.
func Print(t float64) {
	k := Kilogram(t)
	fmt.Printf("%s:", k)
	fmt.Printf("\t%s", KToP(k))

	fmt.Println()

	p := Pound(t)
	fmt.Printf("%s:", p)
	fmt.Printf("\t%s", PToK(p))
}
