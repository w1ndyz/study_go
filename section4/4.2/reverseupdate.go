package main

import "fmt"

func main() {
	a := []byte("Golang大法好")
	b := []rune(string(a))
	fmt.Println(reverse(b))
}

func reverse(arr []rune) string {
	l := len(arr)
	for i := 0; i < l/2; i++ {
		arr[i], arr[l-i-1] = arr[l-i-1], arr[i]
	}
	return string(arr)
}
