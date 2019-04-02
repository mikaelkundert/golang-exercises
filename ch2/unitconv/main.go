// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 43.

// Unitconv converts its numeric argument to Celsius, Fahrenheit and Kelvin.
package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/mikaelkundert/golang-exercises/ch2/lenconv"
	"github.com/mikaelkundert/golang-exercises/ch2/massconv"
	"github.com/mikaelkundert/golang-exercises/ch2/tempconv"
)

func main() {
	if len(os.Args) > 1 {
		for _, arg := range os.Args[1:] {
			t, err := strconv.ParseFloat(arg, 64)
			if err != nil {
				fmt.Fprintf(os.Stderr, "unitconv: %v\n", err)
				os.Exit(1)
			}
			printUnits(t)
		}
	} else {
		input := bufio.NewScanner(os.Stdin)
		for input.Scan() {
			for _, arg := range strings.Split(input.Text(), " ") {
				t, err := strconv.ParseFloat(arg, 64)
				if err != nil {
					fmt.Fprintf(os.Stderr, "unitconv: %v\n", err)
					os.Exit(1)
				}
				printUnits(t)
			}
		}
	}
}

func printUnits(t float64) {
	fmt.Println("TEMPERATURES")
	fmt.Println("============")
	tempconv.Print(t)

	fmt.Println()

	fmt.Println("LENGTHS")
	fmt.Println("=======")
	lenconv.Print(t)

	fmt.Println()

	fmt.Println("MASS")
	fmt.Println("====")
	massconv.Print(t)

	fmt.Println()
}
