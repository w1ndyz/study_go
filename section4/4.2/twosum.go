package main

import "fmt"

func main() {
	var a = []int{1, 2, 3, 5}
	var sum = 4

	b := twoSum(a, sum)
	fmt.Println(b)
}

func twoSum(nums []int, target int) []int {
	result := []int{}
	m := make(map[int]int)
	for i, k := range nums {
		if value, exist := m[target-k]; exist {
			result = append(result, value)
			result = append(result, i)
		}
		m[k] = i
	}
	return result
}
