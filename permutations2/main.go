// https://leetcode.com/problems/permutations-ii
package main

import (
	"fmt"
)

func permute(table map[int]int) [][]int {
	var res [][]int

	for k, count := range table {
		if count <= 0 {
			continue
		}

		table[k] -= 1
		subs := permute(table)
		if len(subs) == 0 {
			res = append(res, []int{k})
		}
		for _, sub := range subs {
			res = append(res, append([]int{k}, sub...))
		}
		table[k] += 1
	}

	return res
}

func permuteUnique(nums []int) [][]int {
	table := make(map[int]int)
	for _, v := range nums {
		table[v] += 1
	}

	return permute(table)
}

func main() {
	fmt.Println(permuteUnique([]int{1, 1, 2, 1}))
}
