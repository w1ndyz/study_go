package main

import "io"

var position int64

type Counting struct {
	W io.Writer
}

func (c *Counting) Write(p []byte) (n int, err error) {
	n, err = c.W.Write(p)
	position += int64(n)
	return n, err
}

func CountingWriter(w io.Writer) (io.Writer, *int64) {
	return &Counting{w}, &position
}
