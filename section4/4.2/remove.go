package main

import "fmt"

func main() {
	a := []string{"a", "a", "b", "c", "d", "d", "e"}
	remove(&a)
	fmt.Println(a)
}

func remove(a *[]string) {
	A := *a
	l := len(A)
	for i := 0; i < l-1; i++ {
		prev := A[i]
		next := A[i+1]
		if prev == next {
			A = append(A[:i], A[i+1:]...)
			l--
		}
	}
	*a = A
}
