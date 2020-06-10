package main

import (
	"bytes"
	"fmt"
)

func (t *tree) String() string {
	var deque []*tree
	var ret []int
	deque = append(deque, t)
	for len(deque) > 0 {
		current := deque[0]
		deque = deque[1:]
		if current.left != nil {
			deque = append(deque, current.left)
		}
		if current.right != nil {
			deque = append(deque, current.right)
		}
		ret = append(ret, current.value)
	}

	var buf bytes.Buffer
	buf.Write([]byte("{"))
	for i, v := range ret {
		if i == len(ret)-1 {
			buf.Write([]byte(fmt.Sprintf("%d", v)))
		} else {
			buf.Write([]byte(fmt.Sprintf("%d, ", v)))
		}
	}
	buf.Write([]byte("}"))
	return buf.String()
}
