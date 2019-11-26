package main

import (
	"fmt"
)

type tuple struct {
	row int
	col int
}

func bfsSudokuSolver(board [][]byte, rows, cols, boxs []int16, spaces []tuple) bool {
	if len(spaces) == 0 {
		return true
	}

	// Get available choices in the first choosable space.
	top := spaces[0]
	row := top.row
	col := top.col
	box := 3*(row/3) + col/3

	var choices []byte
	for _, v := range []byte{0, 1, 2, 3, 4, 5, 6, 7, 8} {
		mask := int16(1 << v)
		if mask&rows[row] > 0 {
			continue
		}
		if mask&cols[col] > 0 {
			continue
		}
		if mask&boxs[box] > 0 {
			continue
		}
		choices = append(choices, v)
	}
	
	if len(choices) == 0 {
		//fmt.Println("no more choices!")
		return false
	}

	for _, v := range choices {
		mask := int16(1 << v)

		// Mask.
		rows[row] |= mask
		cols[col] |= mask
		boxs[box] |= mask

		board[row][col] = byte(v) + byte('1')
		//fmt.Printf("place %c at %+v --- %d len of queue\n", byte(v)+byte('1'), top, len(spaces))
		//print(board)

		ok := bfsSudokuSolver(board, rows, cols, boxs, spaces[1:])
		if ok {
			return true
		}

		board[row][col] = '.'

		// Unmask, unless solution is good!
		rows[row] -= mask
		cols[col] -= mask
		boxs[box] -= mask
	}

	return false
}

func solveSudoku(board [][]byte) {
	rows := make([]int16, 9)
	cols := make([]int16, 9)
	boxs := make([]int16, 9)

	var spaces []tuple

	for row := 0; row < 9; row++ {
		for col := 0; col < 9; col++ {
			box := 3*(row/3) + col/3

			raw := board[row][col]
			if raw == byte('.') {
				spaces = append(spaces, tuple{row, col})
				continue
			}
			val := raw - '1'
			mask := int16(1 << val)
			rows[row] |= mask
			cols[col] |= mask
			boxs[box] |= mask
		}
	}

	bfsSudokuSolver(board, rows, cols, boxs, spaces)
}

func print(board [][]byte) {
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			fmt.Printf("%c ", board[i][j])
		}
		fmt.Println()
	}
}

func main() {
	board := [][]byte{
		{'5', '3', '.', '.', '7', '.', '.', '.', '.'},
		{'6', '.', '.', '1', '9', '5', '.', '.', '.'},
		{'.', '9', '8', '.', '.', '.', '.', '6', '.'},
		{'8', '.', '.', '.', '6', '.', '.', '.', '3'},
		{'4', '.', '.', '8', '.', '3', '.', '.', '1'},
		{'7', '.', '.', '.', '2', '.', '.', '.', '6'},
		{'.', '6', '.', '.', '.', '.', '2', '8', '.'},
		{'.', '.', '.', '4', '1', '9', '.', '.', '5'},
		{'.', '.', '.', '.', '8', '.', '.', '7', '9'},
	}
	solveSudoku(board)
	print(board)
}
