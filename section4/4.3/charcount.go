package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"unicode"
)

func main() {
	var control, digit, graphic, letter, lower, mark, number,
		print, punct, space, symbol, title, upper int

	in := bufio.NewReader(os.Stdin)
	for {
		r, _, err := in.ReadRune()
		if err == io.EOF {
			break
		}
		if unicode.IsControl(r) {
			control++
		}
		if unicode.IsDigit(r) {
			digit++
		}
		if unicode.IsGraphic(r) {
			graphic++
		}
		if unicode.IsLetter(r) {
			letter++
		}
		if unicode.IsLower(r) {
			lower++
		}
		if unicode.IsMark(r) {
			mark++
		}
		if unicode.IsNumber(r) {
			number++
		}
		if unicode.IsPrint(r) {
			print++
		}
		if unicode.IsPunct(r) {
			punct++
		}
		if unicode.IsSpace(r) {
			space++
		}
		if unicode.IsSpace(r) {
			symbol++
		}
		if unicode.IsTitle(r) {
			title++
		}
		if unicode.IsUpper(r) {
			upper++
		}
	}
	fmt.Printf("control: %d\n", control)
	fmt.Printf("digit: %d\n", digit)
	fmt.Printf("graphic: %d\n", graphic)
	fmt.Printf("letter: %d\n", letter)
	fmt.Printf("lower: %d\n", lower)
	fmt.Printf("mark: %d\n", mark)
	fmt.Printf("number: %d\n", number)
	fmt.Printf("print: %d\n", print)
	fmt.Printf("punct: %d\n", punct)
	fmt.Printf("space: %d\n", space)
	fmt.Printf("symbol: %d\n", symbol)
	fmt.Printf("title: %d\n", title)
	fmt.Printf("upper: %d\n", upper)
}
