// https://leetcode.com/problems/max-area-of-island
package main

type tuple struct {
	y, x int
}

func propagate(grid [][]int, coords tuple) int {
	adjacents := []tuple{{1, 0}, {-1, 0}, {0, 1}, {0, -1}}
	valid := func(y, x int) bool {
		return y >= 0 && x >= 0 && y < len(grid) && x < len(grid[0])
	}

	var area int

	queue := []tuple{coords}
	grid[coords.y][coords.x] = 0

	for len(queue) > 0 {
		var next []tuple
		for _, node := range queue {
			area += 1

			for _, delta := range adjacents {
				neighbor := tuple{node.y + delta.y, node.x + delta.x}
				if !valid(neighbor.y, neighbor.x) {
					continue
				}
				if grid[neighbor.y][neighbor.x] == 0 {
					continue
				}

				grid[neighbor.y][neighbor.x] = 0
				next = append(next, neighbor)
			}
		}
		queue = next
	}

	return area
}

func maxAreaOfIsland(grid [][]int) int {
	var best int

	for y, row := range grid {
		for x, cell := range row {
			if cell != 1 {
				continue
			}

			area := propagate(grid, tuple{y, x})
			if area > best {
				best = area
			}
		}
	}

	return best
}
