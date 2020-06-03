package main

import (
	"bufio"
	"bytes"
	"crypto/sha256"
	"fmt"
)

func main() {
	var r1 bytes.Buffer
	writer := bufio.NewWriter(&r1)
	fmt.Fprintf(writer, "%08b", sha256.Sum256([]byte("1")))
	writer.Flush()
	r1Str := r1.String()

	var r2 bytes.Buffer
	writer2 := bufio.NewWriter(&r2)
	fmt.Fprintf(writer2, "%08b", sha256.Sum256([]byte("2")))
	writer2.Flush()
	r2Str := r2.String()

	count := 0
	for i, _ := range r1Str {
		if r1Str[i] != r2Str[i] {
			count += 1
		}
		fmt.Println(count)
	}
}
