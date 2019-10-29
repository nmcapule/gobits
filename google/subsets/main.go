package main

import (
	"fmt"
)

func subsetsOf(set []int) [][]int {
	var result [][]int
	if len(set) <= 1 {
		return result
	}
	for i := 2; i < len(set); i++ {
		result = append(result, set[:i])
	}
	result = append(result, set)
	result = append(result, subsetsOf(set[1:])...)
	return result
}

func main() {
	fmt.Println(subsetsOf([]int{1,2,3,4,5}))
}
