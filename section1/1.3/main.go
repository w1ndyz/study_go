package main

import (
	"bufio"
	"fmt"
	"os"
)

type line struct {
	FileName string
	String   string
}

func main() {
	counts := make(map[line]int)
	files := os.Args[1:]
	if len(files) == 0 {
		countLines(os.Stdin, counts, "ARGS")
	} else {
		for _, arg := range files {
			f, err := os.Open(arg)
			if err != nil {
				fmt.Fprintf(os.Stderr, "dup2: %v\n", err)
				continue
			}
			countLines(f, counts, f.Name())
			f.Close()
		}
	}
	fmt.Println("counts的结构", counts)
	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%s文件\t%s字符串\t出现了%d次\n", line.FileName, line.String, n)
		}
	}
}

func countLines(f *os.File, counts map[line]int, fileName string) {
	input := bufio.NewScanner(f)
	for input.Scan() {
		counts[line{
			FileName: fileName,
			String:   input.Text(),
		}]++
	}
}
