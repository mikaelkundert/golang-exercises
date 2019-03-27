package main

import (
	"fmt"
	"os"
)

func main() {
	var delimiter = ": "
	for index, arg := range os.Args[1:] {
		fmt.Print(index + 1)
		fmt.Println(delimiter + arg)
	}
}
