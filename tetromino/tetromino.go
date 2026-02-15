package tetromino

type Block struct {
	X, Y int
}


type Tetromino struct {
	Blocks [4]Block
	Letter rune
}

//var pieces []Tetromino