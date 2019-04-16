package main

import (
	"bytes"
	"fmt"
	"os"
	"strings"
)

func main() {
	if len(os.Args) <= 2 {
		fmt.Fprintf(os.Stderr, "Error: Too few arguments.\n\nUsage:\n  %s value1 value2 [valueN, ...]\n", os.Args[0])
		return
	}
	var anagram = true
	for i := 2; i < len(os.Args); i++ {
		if !isAnagram(os.Args[i-1], os.Args[i]) {
			anagram = false
			break
		}
	}
	if anagram {
		fmt.Println("Given words are anagram.")
	} else {
		fmt.Println("Given words are not anagram.")
	}
}

func isAnagram(s1 string, s2 string) bool {
	if len(s1) != len(s2) {
		return false
	}
	s2 = strings.ToLower(s2)
	for i := 0; i < len(s1); i++ {
		if index := strings.Index(s2, s1[i:i+1]); index != -1 {
			var buf bytes.Buffer
			for b := 0; b < len(s2); b++ {
				if b != index {
					if _, err := buf.WriteString(s2[b : b+1]); err != nil {
						panic(fmt.Sprintf("Could not write new cut string."))
					}
				}
			}
			s2 = buf.String()
		}
	}
	return len(s2) == 0
}
