package main

import (
	"crypto/sha256"
	"crypto/sha512"
	"flag"
	"fmt"
)

func main() {
	method := flag.String("method", "sha256", "select hash method(sha256,sha384,sha512)")
	text := flag.String("text", "", "input the string you want to hash")
	flag.Parse()
	switch *method {
	case "sha256":
		fmt.Printf("%x\n", sha256.Sum256([]byte(*text)))
	case "sha384":
		fmt.Printf("%x\n", sha512.Sum384([]byte(*text)))
	case "sha512":
		fmt.Printf("%x\n", sha512.Sum512([]byte(*text)))
	default:
		fmt.Printf("not support")
	}
}
