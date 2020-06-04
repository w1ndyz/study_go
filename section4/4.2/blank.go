package main

import (
	"fmt"
	"unicode"
)

func main() {
	a := "asd asdasd kjhiuweyr   ar"
	removeEmpty(&a)
	fmt.Println(a)
}

func removeEmpty(s *string) {
	S := *s
	str := string(S[0])
	end := 0
	for _, v := range S {
		last := rune(str[end])
		if unicode.IsSpace(last) && unicode.IsSpace(v) {
			continue
		}
		str += string(v)
		end++
	}
	fmt.Println(str)
	*s = str[1:]
}
