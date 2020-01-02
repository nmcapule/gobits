// https://leetcode.com/problems/combination-sum-ii
package main

import (
	"fmt"
	"sort"
)

type tuple struct {
	value int
	count int
}

func recursiveK2(pairs []*tuple, target int) [][]int {
	if len(pairs) == 0 {
		return nil
	}

	head := pairs[0]
	if head.count < 1 {
		return recursiveK2(pairs[1:], target)
	}
	if head.value == target {
		return [][]int{{head.value}}
	}

	var res [][]int
	if head.value <= target {
		head.count -= 1
		subs := recursiveK2(pairs, target-head.value)
		head.count += 1
		for _, sub := range subs {
			res = append(res, append([]int{head.value}, sub...))
		}
	}
	head.count -= 1
	subs := recursiveK2(pairs[1:], target)
	head.count += 1
	for _, sub := range subs {
		res = append(res, sub)
	}

	return res
}

func combinationSum2(candidates []int, target int) [][]int {
	sort.Ints(candidates)

	var pairs []*tuple
	for _, value := range candidates {
		if len(pairs) == 0 {
			pairs = append(pairs, &tuple{value, 1})
			continue
		}
		tail := pairs[len(pairs)-1]
		if tail.value != value {
			pairs = append(pairs, &tuple{value, 1})
			continue
		}
		tail.count += 1
	}

	return recursiveK2(pairs, target)
}

func main() {
	fmt.Println(combinationSum2([]int{2,5,2,1,2}, 5))
}
