// https://leetcode.com/problems/n-queens/
package main

import (
	"bytes"
	"fmt"
)

func recursiveSolve(n, row int, verticals, lrdiagonals, rldiagonals []bool) [][]int {
	if row == n {
		return nil
	}

	var results [][]int
	for col := 0; col < n; col++ {
		vidx := col
		lridx := col - row + n - 1
		rlidx := col + row

		if verticals[vidx] {
			continue
		}
		if lrdiagonals[lridx] {
			continue
		}
		if rldiagonals[rlidx] {
			continue
		}

		verticals[vidx] = true
		lrdiagonals[lridx] = true
		rldiagonals[rlidx] = true

		if row == n-1 {
			results = append(results, []int{col})
		} else {
			subs := recursiveSolve(n, row+1, verticals, lrdiagonals, rldiagonals)
			for _, sub := range subs {
				results = append(results, append([]int{col}, sub...))
			}
		}

		verticals[vidx] = false
		lrdiagonals[lridx] = false
		rldiagonals[rlidx] = false
	}
	return results
}

func main() {
	n := 10
	verticals := make([]bool, n)
	lrdiagonals := make([]bool, n*2-1)
	rldiagonals := make([]bool, n*2-1)

	solutions := recursiveSolve(n, 0, verticals, lrdiagonals, rldiagonals)
	
	lookup := make([]string, n)
	for i := 0; i < n; i++ {
		var buf bytes.Buffer
		for j := 0; j < n; j++ {
			buf.WriteByte(' ')
			if j == i {
				buf.WriteByte('Q')
				continue
			}
			buf.WriteByte('.')
		}
		lookup[i] = buf.String()
	}
	
	var results [][]string
	for _, idxs := range solutions {
		var lines []string
		for _, pos := range idxs {
			lines = append(lines, lookup[pos])
			fmt.Println(lookup[pos])
		}
		fmt.Println("-----")
		results = append(results, lines)
	}
	fmt.Println(results)
}
