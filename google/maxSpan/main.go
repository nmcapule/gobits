package main

import "fmt"

type tuple struct {
	left  int
	right int
}

func (t *tuple) span() int {
	return t.right - t.left + 1
}

func maxSpan(arr []int) int {
	lookup := make(map[int]*tuple)
	for i, v := range arr {
		if _, ok := lookup[v]; !ok {
			lookup[v] = &tuple{i, i}
		}

		lookup[v].right = i
	}

	var best int
	for _, t := range lookup {
		if best < t.span() {
			best = t.span()
		}
	}
	return best
}

func main() {
	fmt.Println(maxSpan([]int{1, 2, 1, 1, 3}))
	fmt.Println(maxSpan([]int{1, 4, 2, 1, 4, 1, 4}))
	fmt.Println(maxSpan([]int{1, 4, 2, 1, 4, 4, 4}))
}
