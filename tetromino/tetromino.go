package tetromino

import "fmt"

type Block struct {
	X, Y int
}

type Tetromino struct {
	Blocks [4]Block
	Letter rune
}

//var pieces []Tetromino

// normalizeblocks to start at (0,0)
func BuildNewTetromino(lines []string, letter rune) (Tetromino, error) {
	if len(lines) != 4 {
		return Tetromino{}, fmt.Errorf("invalid piece")
	}

	var blocks []Block
	for y := 0; y < 4; y++ {
		if len(lines[y]) != 4 {
			return Tetromino{}, fmt.Errorf("invalid line length")
		}
		for x := 0; x < 4; x++ {
			if lines[y][x] == '#' {
				blocks = append(blocks, Block{X: x, Y: y})
			} else if lines[y][x] != '.' {
				return Tetromino{}, fmt.Errorf("invalid character")

			}
		}
	}
	if len(blocks) != 4 {
		return Tetromino{}, fmt.Errorf("invalid number of blocks")
	}

	if !isConnected(blocks) {
		return Tetromino{}, fmt.Errorf("blocks not connected")
	}

	//normalize to 0,0
	minX, minY := blocks[0].X, blocks[0].Y
	for _, b := range blocks {
		if b.X < minX {
			minX = b.X
		}
		if b.Y < minY {
			minY = b.Y
		}
	}

	var t Tetromino
	t.Letter = letter
	for i := 0; i < 4; i++ {
		t.Blocks[i] = Block{
			X: blocks[i].X - minX,
			Y: blocks[i].Y - minY,
		}
	}
	return t, nil
}

// check if shape is validly connected
func isConnected(blocks []Block) bool {
	touches := 0

	for i := 0; i < len(blocks); i++ {
		for j := i + 1; j < len(blocks); j++ {
			if i == j {
				continue
			}

			dx := blocks[i].X - blocks[j].X
			if dx < 0 {
				dx = -dx
			}
			dy := blocks[i].Y - blocks[j].Y
			if dy < 0 {
				dy = -dy
			}

			// directly adjacent
			if (dx == 1 && dy == 0) || (dx == 0 && dy == 1) {
				touches++
			}
		}
	}

	// valid tetrominoes have at least 6 side connections
	return touches >= 3
}
