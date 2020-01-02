package main

import "fmt"

// O(N)
func canBalance(input []int) bool {
	var left, right int

	// Pool all to right.
	for _, x := range input {
		right += x
	}

	// Incrementally check sum for left, then compare.
	for _, x := range input {
		left += x
		right -= x
		if left == right {
			return true
		}
	}

	return false
}

func main() {
	fmt.Println(canBalance([]int{1, 1, 1, 2, 1}))
	fmt.Println(canBalance([]int{2, 1, 1, 2, 1}))
	fmt.Println(canBalance([]int{10, 10}))
}
