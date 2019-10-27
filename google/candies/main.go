// https://leetcode.com/problems/distribute-candies/
package main

import "fmt"

func distributeCandies(candies []int) int {
	lookup := make(map[int]bool)
	for _, c := range candies {
		lookup[c] = true
	}

	if len(lookup) > len(candies)/2 {
		return len(candies) / 2
	}
	return len(lookup)
}

func main() {
	fmt.Println(distributeCandies([]int{1, 1, 2, 3, 2, 3}))
	fmt.Println(distributeCandies([]int{1, 1, 2, 3}))
}
