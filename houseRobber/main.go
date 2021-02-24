package main

import "fmt"

func rob(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	if len(nums) == 1 {
		return nums[0]
	}

	for i := 2; i < len(nums); i++ {
		var a, b int
		if i-2 >= 0 {
			a = nums[i] + nums[i-2]
		}
		if i-3 >= 0 {
			b = nums[i] + nums[i-3]
		}

		if a > b {
			nums[i] = a
		} else {
			nums[i] = b
		}
	}

	if nums[len(nums)-1] > nums[len(nums)-2] {
		return nums[len(nums)-1]
	}
	return nums[len(nums)-2]
}

func main() {
	fmt.Println(rob([]int{2, 7, 9, 3, 1}))
}
