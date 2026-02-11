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


