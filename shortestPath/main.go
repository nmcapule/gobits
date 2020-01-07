// Daily coding #23
//
// You are given an M by N matrix consisting of booleans that represents a board.
// Each True boolean represents a wall. Each False boolean represents a tile you
// can walk on.
//
// Given this matrix, a start coordinate, and an end coordinate, return the minimum
// number of steps required to reach the end coordinate from the start. If there
// is no possible path, then return null. You can move up, left, down, and right.
// You cannot move through walls. You cannot wrap around the edges of the board.
//
// For example, given the following board:
//
// [[f, f, f, f],
// [t, t, f, t],
// [f, f, f, f],
// [f, f, f, f]]
//
// and start = (3, 0) (bottom left) and end = (0, 0) (top left), the minimum number
// of steps required to reach the end is 7, since we would need to go through (1, 2)
// because there is a wall everywhere else on the second row.
package main

import (
	"fmt"
)

type coords struct {
	y, x int
}

func shortestPathSteps(mat [][]bool, start, end coords) (int, bool) {
	valid := func(i, j int) bool {
		return i >= 0 && j >= 0 && i < len(mat) && j < len(mat[0])
	}
	adjacents := []coords{{1, 0}, {-1, 0}, {0, 1}, {0, -1}}

	costs := make(map[coords]int)

	queue := []coords{start}
	costs[start] = 0

	for len(queue) > 0 {
		fmt.Println(queue)
		var next []coords
		for _, node := range queue {
			for _, delta := range adjacents {
				neighbor := coords{node.y + delta.y, node.x + delta.x}
				if !valid(neighbor.y, neighbor.x) {
					continue
				}
				if mat[neighbor.y][neighbor.x] == true {
					continue
				}
				cost, ok := costs[neighbor]
				if !ok || cost > costs[node]+1 {
					next = append(next, neighbor)
					costs[neighbor] = costs[node] + 1
				}
			}
		}
		queue = next
	}

	cost, ok := costs[end]
	return cost, ok
}

func main() {
	mat := [][]bool{
		{false, false, false, false, false},
		{true, true, true, true, false},
		{false, false, false, true, false},
		{false, true, false, false, false}}

	start := coords{3, 0}
	end := coords{0, 0}

	fmt.Println(shortestPathSteps(mat, start, end))
}
