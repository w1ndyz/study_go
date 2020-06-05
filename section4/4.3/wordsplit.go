package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	scanner := bufio.NewScanner(bufio.NewReader(os.Stdin))
	scanner.Split(bufio.ScanWords)
	record := map[string]int{}

	for scanner.Scan() {
		record[scanner.Text()]++
	}

	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "Reading input", err)
	}

	for k, v := range record {
		fmt.Printf("%s\t%d\n", k, v)
	}
}
