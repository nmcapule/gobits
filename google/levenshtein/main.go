// https://leetcode.com/problems/simplify-path/
package main

import (
	"fmt"
)

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func minDistance(word1 string, word2 string) int {
	idx := func(i, j int) int {
		return i*(len(word2)+1) + (j % (len(word2) + 1))
	}

	mat := make([]int, (len(word1)+1)*(len(word2)+1))
	for i := 0; i <= len(word1); i++ {
		for j := 0; j <= len(word2); j++ {
			if i == 0 || j == 0 {
				mat[idx(i, j)] = max(i, j)
				continue
			}

			var mismatchCost int
			if word1[i-1] != word2[j-1] {
				mismatchCost = 1
			}

			mat[idx(i, j)] = min(
				mat[idx(i-1, j-1)]+mismatchCost,
				min(mat[idx(i-1, j)], mat[idx(i, j-1)])+1,
			)
		}
	}

	return mat[len(mat)-1]
}

func main() {
	fmt.Println(minDistance("23", "1"))
	fmt.Println(minDistance("horse", "ros"))
	fmt.Println(minDistance("ros", "horse"))
	fmt.Println(minDistance("intention", "execution"))
	fmt.Println(minDistance("execution", "intention"))
}
