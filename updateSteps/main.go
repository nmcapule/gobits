// Given a bool matrix, how many steps will it take for all cells to be marked
// "true".
//
// On every step, cells adjacent to "true" will be marked as "true".
package main

import (
	"fmt"
)

type coords struct {
	y, x int
}

func propagateUpdateSteps(mat [][]bool) int {
	valid := func(i, j int) bool {
		return i >= 0 && j >= 0 && i < len(mat) && j < len(mat[0])
	}
	adjacents := []coords{{1, 0}, {-1, 0}, {0, 1}, {0, -1}}

	virals := make(map[coords]bool)
	zeroes := make(map[coords]bool)

	for y, row := range mat {
		for x, v := range row {
			if v {
				virals[coords{y, x}] = true
			} else {
				zeroes[coords{y, x}] = true
			}
		}
	}

	var steps int
	for len(zeroes) > 0 && steps < len(mat) * len(mat[0]) {
		next := make(map[coords]bool)
		for node := range virals {
			for _, delta := range adjacents {
				neighbor := coords{node.y + delta.y, node.x + delta.x}
				if !valid(neighbor.y, neighbor.x) {
					continue
				}
				if mat[neighbor.y][neighbor.x] == true {
					continue
				}
				_, ok := zeroes[neighbor]
				if ok {
					next[neighbor] = true
					delete(zeroes, neighbor)
				}
			}
		}
		virals = next

		steps += 1
	}

	return steps
}

func main() {
	mat := [][]bool{
		{false, false, false, false, false},
		{false, false, true, false, false},
		{false, false, false, false, false}}

	fmt.Println(propagateUpdateSteps(mat))
}
