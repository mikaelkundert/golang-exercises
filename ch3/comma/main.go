// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 73.

// Comma prints its argument numbers with a comma at each power of 1000.
//
// Example:
// 	$ go build gopl.io/ch3/comma
//	$ ./comma 1 12 123 1234 1234567890
// 	1
// 	12
// 	123
// 	1,234
// 	1,234,567,890
//
package main

import (
	"bytes"
	"fmt"
	"os"
)

func main() {
	for i := 1; i < len(os.Args); i++ {
		fmt.Printf("  %s\n", comma(os.Args[i]))
	}
}

//!+
// comma inserts commas in a non-negative decimal integer string.
func comma(s string) string {
	var buf bytes.Buffer
	var tens int8
	for i := len(s); i > 0; i-- {
		buf.WriteString(s[i-1 : i])
		if tens < 2 {
			tens++
		} else {
			if i != 1 {
				buf.WriteString(",")
			}
			tens = 0
		}
	}
	return reverse(buf.String())
}

func reverse(s string) string {
	var buf bytes.Buffer
	for i := len(s); i > 0; i-- {
		buf.WriteString(s[i-1 : i])
	}
	return buf.String()
}

//!-
