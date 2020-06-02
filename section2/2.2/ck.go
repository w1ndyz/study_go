package main

import (
	"fmt"
	"os"
	"strconv"

	"study_go/section2/2.1/tempconv"
)

func main() {
	for _, arg := range os.Args[1:] {
		t, err := strconv.ParseFloat(arg, 64)
		if err != nil {
			fmt.Fprintf(os.Stderr, "cf: %v\n", err)
			os.Exit(1)
		}
		c := tempconv.Celsius(t)
		fmt.Printf("%s = %.2f\n", c, tempconv.CToK(c))
	}
}
