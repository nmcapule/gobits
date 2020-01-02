// https://leetcode.com/problems/first-missing-positive
package main

import (
	"fmt"
)

func firstMissingPositive(nums []int) int {
	// Sanitize. Remove negatives and zero.
	// O(n)
	var slot int
	for i := 0; i < len(nums); i++ {
		if nums[i] > 0 {
			nums[slot] = nums[i]
			slot += 1
		}
	}
	nums = nums[:slot]

	if len(nums) == 0 {
		return 1
	}

	// Check min.
	// O(n)
	min := nums[0]
	for i := 1; i < len(nums); i++ {
		if nums[i] < min {
			min = nums[i]
		}
	}
	if min > 1 {
		return 1
	}

	// Inplace bucket sort!
	// Should be O(n)
	var i int
	for i < len(nums) {
		num := nums[i]
		pos := num - 1

		// If current num is empty (-1), move on.
		if num == -1 {
			i += 1
			continue
		}
		// If current num fits its position, move on.
		if pos == i {
			i += 1
			continue
		}
		// If current num's proper position exceeds len, set cur to -1 and move on.
		if pos >= len(nums) {
			nums[i] = -1
			i += 1
			continue
		}

		// If current num = target pos num, set current to -1.
		if nums[i] == nums[pos] {
			nums[i] = -1
			i += 1
			continue
		}

		// Else, swap to its proper position.
		tmp := nums[pos]
		nums[pos] = nums[i]
		nums[i] = tmp
	}

	for i := 0; i < len(nums); i++ {
		if nums[i] == -1 {
			return i + 1
		}
	}
	return len(nums) + 1
}

func main() {
	fmt.Println(firstMissingPositive([]int{1, 1}))
}
