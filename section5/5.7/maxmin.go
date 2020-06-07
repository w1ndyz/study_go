package main

import "math"

func Max(n ...int) int {
	if len(n) == 0 {
		panic("at least one number")
	}
	ret := math.MinInt64
	for _, v := range n {
		if v > ret {
			ret = v
		}
	}
	return ret
}

func Min(n ...int) int {
	if len(n) == 0 {
		panic("at least one number")
	}
	ret := math.MaxInt64
	for _, v := range n {
		if v < ret {
			ret = v
		}
	}
	return ret
}
