package main

import (
	"bufio"
	"crypto/sha256"
	"fmt"
	"os"
)

func main() {
	var scanner = bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		hash := sha256.Sum256(scanner.Bytes())
		fmt.Printf("%x\n", string(hash[:]))
	}
}
