package main

import "fmt"

func main() {
	s := []int{0, 1, 2, 3, 4, 5, 6, 7}
	r := rotate(s, 3)
	fmt.Println(r)
}

func rotate(s []int, position int) []int {
	r := s[position:]
	for i := position - 1; i >= 0; i-- {
		r = append(r, s[i])
	}
	return r
}
