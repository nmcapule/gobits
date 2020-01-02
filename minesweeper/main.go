// https://techdevguide.withgoogle.com/paths/foundational/coding-question-minesweeper/#!
package main

import (
	"bytes"
	"fmt"
	"math/rand"
	"time"
)

type board struct {
	rows, cols int
	matrix     []rune
}

func newBoard(rows, cols, mines int) *board {
	b := &board{
		rows:   rows,
		cols:   cols,
		matrix: make([]rune, rows*cols),
	}

	for i := 0; i < len(b.matrix); i++ {
		if i < mines {
			b.matrix[i] = rune('x')
		} else {
			b.matrix[i] = rune('_')
		}
	}

	// Shuffle our board.
	for i := len(b.matrix) - 1; i > 0; i-- {
		victim := rand.Intn(i)

		tmp := b.matrix[victim]
		b.matrix[victim] = b.matrix[i]
		b.matrix[i] = tmp
	}

	return b
}

func (b board) String() string {
	var buffer bytes.Buffer
	for i, r := range b.matrix {
		buffer.WriteRune(r)
		buffer.WriteRune(' ')
		if (i+1)%b.cols == 0 {
			buffer.WriteRune('\n')
		}
	}
	return buffer.String()
}

func main() {
	rand.Seed(time.Now().UTC().UnixNano())

	b := newBoard(30, 50, 1000)
	fmt.Println(b.String())
}
