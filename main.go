package main

type Block struct {
	x, y int
}

type Tetromino struct {
	blocks [4]Block
	letter rune
}

var pieces []Tetromino

type Board struct {
	cells [][]rune
	size  int
}

//Board methods first
/*r = row + block.y → board row
c = col + block.x → board column
Loop through all 4 blocks of the piece
If any block is out of bounds or overlaps → cannot place*/

func (b *Board) CanPlace(t Tetromino, row, col int) bool {
	for _, block := range t.blocks {
		r := row + block.y
		c := col + block.x

		if r < 0 || r >= b.size || c >= b.size {
			return false
		}

		if b.cells[r][c] != '.' {
			return false
		}
	}
	return true
}

// actually place
func (b *Board) Place(t Tetromino, row, col int) {
	for _, block := range t.blocks {
		b.cells[row+block.y][col+block.x] = t.letter
	}
}

// remove for backtracking
func (b *Board) remove(t Tetromino, row, col int) {
	for _, block := range t.blocks {
		b.cells[row+block.y][col+block.x] = '.'
	}
}

// loop over the pieces
func Solve(b *Board, pieces []Tetromino, index int) bool {

	if index == len(pieces) {
		return false
	}

	t := pieces[index]

	for row := 0; row < b.size; row++ {
		for col := 0; col < b.size; col++ {

			if b.CanPlace(t, row, col) {
				b.Place(t, row, col)

				if Solve(b, pieces, index+1) {
					return true
				}

				b.remove(t, row, col)
			}
		}
	}
	return false
}
