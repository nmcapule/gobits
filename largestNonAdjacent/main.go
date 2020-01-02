// Good morning! Here's your coding interview problem for today.
// This problem was asked by Airbnb.
// Given a list of integers, write a function that returns the largest sum of non-adjacent numbers. Numbers can be 0 or negative.
// For example, [2, 4, 6, 2, 5] should return 13, since we pick 2, 6, and 5. [5, 1, 1, 5] should return 10, since we pick 5 and 5.
// Follow-up: Can you do this in O(N) time and constant space?
package main

import (
	"fmt"
)

func largestNonAdjacent(nums []int) int {
	var best int
	for i, _ := range nums {
		if i == 2 {
			nums[i] = nums[i] + nums[i-2]
		}
		if i >= 3 {
			candidate1 := nums[i] + nums[i-2]
			candidate2 := nums[i] + nums[i-3]
			if candidate1 > candidate2 {
				nums[i] = candidate1
			} else {
				nums[i] = candidate2
			}
		}
		if nums[i] > best {
			best = nums[i]
		}
	}
	return best
}

func main() {
	fmt.Println(largestNonAdjacent([]int{2,4,3,3,5,4,4,5,1}))
	fmt.Println(largestNonAdjacent([]int{5,1,1,5}))
}
