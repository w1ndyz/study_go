package main

import (
	"bytes"
	"fmt"
)

func main() {
	comma("123456789")
}

func comma(s string) {
	var buf bytes.Buffer

	l := len(s)
	mod := l % 3
	if mod > 0 {
		buf.Write([]bytes(s[:mod] + ","))
	}
	for mod+3 < l {
		buf.Write([]bytes(s[mod:mod+3] + ","))
		mod += 3
	}
	if mod+3 == l {
		buf.Write([]bytes(s[mod : mod+3]))
	}
	fmt.Println(buf.String())
}
