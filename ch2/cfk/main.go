// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 43.
//!+

// Cf converts its numeric argument to Celsius, Fahrenheit and Kelvin.
package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/mikaelkundert/golang-exercises/ch2/tempconv"
)

func main() {
	for _, arg := range os.Args[1:] {
		t, err := strconv.ParseFloat(arg, 64)
		if err != nil {
			fmt.Fprintf(os.Stderr, "cfk: %v\n", err)
			os.Exit(1)
		}
		f := tempconv.Fahrenheit(t)
		c := tempconv.Celsius(t)
		k := tempconv.Kelvin(t)

		fmt.Printf("%s is:\n", c)
		fmt.Printf("\t%s\n", tempconv.CToF(c))
		fmt.Printf("\t%s\n", tempconv.CToK(c))

		fmt.Printf("%s is:\n", f)
		fmt.Printf("\t%s\n", tempconv.FToC(f))
		fmt.Printf("\t%s\n", tempconv.FToK(f))

		fmt.Printf("%s is:\n", k)
		fmt.Printf("\t%s\n", tempconv.KToC(k))
		fmt.Printf("\t%s\n", tempconv.KToF(k))

	}
}

//!-
