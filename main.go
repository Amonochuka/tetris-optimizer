package main

import "tetris-optimizer/board"


//newboard constructor

func NewBoard(size int) *board.Board{
	cells := make([][]rune, size)
	for i := range cells{
		cells[i] = make([]rune, size)
		for j := range cells[i]{
			cells[i][j] = '.'
		}
	}

	return &board.Board{
		Cells: cells,
		Size: size,
	}
}
