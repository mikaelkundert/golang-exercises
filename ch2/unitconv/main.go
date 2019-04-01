// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 43.
//!+

// Unitconv converts its numeric argument to Celsius, Fahrenheit and Kelvin.
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
			fmt.Fprintf(os.Stderr, "unitconv: %v\n", err)
			os.Exit(1)
		}
		tempconv.Print(t)
	}
}

//!-
