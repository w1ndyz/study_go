package main

import "fmt"

func main() {
	var s = []int{1, 2, 3}
	reverse(&s)
	fmt.Println(s)
}

func reverse(s *[]int) {
	A := *s
	l := len(A)
	for i := 0; i < l/2; i++ {
		A[i], A[l-i-1] = A[l-i-1], A[i]
	}
}
