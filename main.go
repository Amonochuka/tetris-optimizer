package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"tetris-optimizer/board"
	"tetris-optimizer/solver"
	"tetris-optimizer/tetromino"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("ERROR")
		return
	}

	file, err := os.Open(os.Args[1])
	if err != nil {
		fmt.Println("ERROR")
		return
	}
	defer file.Close()

	pieces, err := readPieces(file)
	if err != nil {
		fmt.Println("ERROR")
		return
	}

	//check if at least one piece is present
	if len(pieces) == 0 {
		fmt.Println("ERROR")
		return
	}

	// for now, just print how many pieces we got
	size := minBoardSize(len(pieces))

	for {
		b := board.NewBoard(size)

		if solver.Solve(&b, pieces, 0) {
			printBoard(&b)
			return
		}

		size++
	}

}

func readPieces(file *os.File) ([]tetromino.Tetromino, error) {
	scanner := bufio.NewScanner(file)

	var pieces []tetromino.Tetromino
	var blockLines []string
	letter := 'A'

	for scanner.Scan() {
		line := strings.TrimRight(scanner.Text(), "\r")

		if line == "" {
			if len(blockLines) != 0 {
				t, err := tetromino.BuildNewTetromino(blockLines, letter)
				if err != nil {
					return nil, err
				}
				pieces = append(pieces, t)
				letter++
				blockLines = nil
			}
			continue
		}

		blockLines = append(blockLines, line)
	}

	// handle last piece if file ends without empty line
	if len(blockLines) != 0 {
		t, err := tetromino.BuildNewTetromino(blockLines, letter)
		if err != nil {
			return nil, err
		}
		pieces = append(pieces, t)
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return pieces, nil

	//testtingif file open works before adding other parts
	//fmt.Println("OK")
}

func minBoardSize(pieceCount int) int {
	area := pieceCount * 4
	size := 2

	for size*size < area {
		size++
	}

	return size
}

func printBoard(b *board.Board) {
	for i := 0; i < b.Size; i++ {
		for j := 0; j < b.Size; j++ {
			fmt.Print(string(b.Cells[i][j]))
		}
		fmt.Println()
	}
}
