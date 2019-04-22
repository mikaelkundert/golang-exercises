package main

import (
	"bufio"
	"bytes"
	"crypto/sha256"
	"crypto/sha512"
	"flag"
	"fmt"
	"os"
)

var hash = flag.String("h", "SHA256", "Specify hash type (SHA256, SHA384 or SHA512). Default: SHA256")

func main() {
	flag.Parse()
	var scanner = bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		var input = scanner.Bytes()
		var buf bytes.Buffer
		switch *hash {
		case "SHA256":
			for _, b := range sha256.Sum256(input) {
				buf.WriteByte(b)
			}
		case "SHA384":
			for _, b := range sha512.Sum384(input) {
				buf.WriteByte(b)
			}
		case "SHA512":
			for _, b := range sha512.Sum512(input) {
				buf.WriteByte(b)
			}
		}
		if len(buf.String()) > 0 {
			fmt.Printf("%x\n", buf.String())
		}
	}
}
