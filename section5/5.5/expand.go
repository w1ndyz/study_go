package main

import "strings"

func Expand(s string, f func(string) string) string {
	ret := strings.Replace(s, "$foo", f("foo"), 1024)
	return ret
}
