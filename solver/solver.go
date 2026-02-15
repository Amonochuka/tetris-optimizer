package solver

import ("tetris-optimizer/tetromino"
"tetris-optimizer/board")

// loop over the pieces
func Solve(b *board.Board, pieces []tetromino.Tetromino, index int) bool {

	if index == len(pieces) {
		return true
	}

	t := pieces[index]

	for row := 0; row < b.Size; row++ {
		for col := 0; col < b.Size; col++ {

			if b.CanPlace(t, row, col) {
				b.Place(t, row, col)

				if Solve(b, pieces, index+1) {
					return true
				}

				b.Remove(t, row, col)
			}
		}
	}
	return false
}
