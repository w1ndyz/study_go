package main

import (
	"bufio"
	"fmt"
	"strings"
)

type ByteCounter int

func (c *ByteCounter) Write(p []byte) (int, error) {
	scan := bufio.NewScanner(strings.NewReader(string(p)))
	scan.Split(bufio.ScanWords)

	for scan.Scan() {
		*c += 1
	}
	return len(p), nil
}

func main() {
	var c ByteCounter

	var name = "Windy rain"
	fmt.Fprintf(&c, "hello, %s", name)
	fmt.Println(c)
}
