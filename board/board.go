package board

import "tetris-optimizer/tetromino"

type Board struct {
	Cells [][]rune
	Size  int
}

//Board methods first
/*r = row + block.y → board row
c = col + block.x → board column
Loop through all 4 blocks of the piece
If any block is out of bounds or overlaps → cannot place*/

func (b *Board) CanPlace(t tetromino.Tetromino, row, col int) bool {
	for _, block := range t.Blocks {
		r := row + block.Y
		c := col + block.X

		if r < 0 || r >= b.Size || c >= b.Size {
			return false
		}

		if b.Cells[r][c] != '.' {
			return false
		}
	}
	return true
}

// actually place
func (b *Board) Place(t tetromino.Tetromino, row, col int) {
	for _, block := range t.Blocks {
		b.Cells[row+block.Y][col+block.X] = t.Letter
	}
}

// remove for backtracking
func (b *Board) Remove(t tetromino.Tetromino, row, col int) {
	for _, block := range t.Blocks {
		b.Cells[row+block.Y][col+block.X] = '.'
	}
}
