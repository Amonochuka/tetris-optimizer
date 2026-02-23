
# Tetris-optimizer

A Go program that reads a list of tetrominoes from a text file and assembles them into the **smallest possible square**.

Each tetromino is printed using uppercase letters (`A`, `B`, `C`, …) in the order they appear in the file.

## Requirements

- Go (standard library only)
- One argument: path to the input file

## How to run

``` bash
go run . sample.txt
```

## Input format

- Each tetromino:

  - is exactly 4 lines

  - each line is exactly 4 characters

  - allowed characters: '.' and '#'

  - tetrominoes are separated by an empty line

  - the file may or may not end with a blank line

- Example:
```
#...
#...
#...
#...
```


## Rules

  - Each tetromino must contain exactly 4 # blocks

  - All '#' blocks must be connected by sides

**Any invalid format must print:**
ERROR

At least one tetromino is expected

The program tries to fit all pieces into the smallest square possible

If a perfect square is not possible, empty cells are left as '.'

Example

Input:
```

...#
...#
...#
...#

....
....
....
####

.###
...#
....
....
```

```Run:
go run . sample.txt
```

Output:

```
ABBB.
ACCC.
A....
A....
```

(letters may vary depending on the input order)

## Project structure

```
├── main.go
├── board/
│   └── board.go
├── solver/
│   └── solver.go
└── tetromino/
    └── tetromino.go
```


## Tests

Basic unit tests are provided for tetromino validation.

Run all tests with:

go test ./...

## Notes

  - The program uses a backtracking solver.

  - Tetrominoes are normalized so their coordinates always start at (0,0).

  - Only the Go standard library is used.

## How the Code works; a breakdown

## Step-by-step algorithm

***Step 1 — Parse the input file***

 - Use bufio.Scanner to read the file line by line

 - Each tetromino is 4 lines of 4 characters (enforced in parseTetromino)

 - A blank line separates tetrominoes

```
Each tetromino is stored in a Tetromino struct:

type Tetromino struct {
	Blocks [4]Block
	Letter rune
}
```

 - Each Block has coordinates {X, Y} relative to the 4×4 grid

 - Normalization shifts tetrominoes to the top-left for easier placement later

Why this approach:

 - Reading line-by-line with a slice lines []string matches the text file format exactly

 - Normalizing ensures placement logic doesn’t have to deal with arbitrary offsets

***Step 2 — Validate Tetromino***

In parseTetromino:

 - Ensure there are exactly 4 lines and 4 # blocks
 
 - Only '.' and '#' are allowed characters

 - Store coordinates of blocks in an array [4]Block

Why this approach:

 - Enforces input format strictly

 - Fixed-size array [4] avoids dynamic allocation for known-size data

 - Makes solver logic simpler (we know every tetromino has exactly 4 blocks)

***Step 3 — Minimum board size calculation***

```
size := 1
for size*size < n*4 {
    size++
}

Where n = len(tetromino.Pieces)
```

***Explanation:***

 - Each tetromino occupies 4 cells

 - Minimal square size = smallest size such that size*size >= total_cells

 - Avoids wasting space and ensures the algorithm starts with the smallest possible board

***Step 4 — Solve with Backtracking***

solver.Solve:

 - Recursive backtracking approach:

 - Try to place the current tetromino at every possible position on the board

 - If placement succeeds, recurse to the next tetromino

  - If stuck, remove the last piece (backtrack)

Base case: all pieces placed → return true

Why this approach:

 - Backtracking guarantees that all arrangements are explored

 - Works efficiently because:

 - Board size starts minimal

 - Pieces are normalized

 - Only valid positions are attempted (CanPlace)

***Step 5 — Board operations***

board.Board has methods:

 - CanPlace(t, row, col) → checks if a tetromino fits

 - Place(t, row, col) → marks the board with the tetromino letter
 
 - Remove(t, row, col) → clears the board (for backtracking)

Why methods instead of functions:

 - Encapsulation: board knows its own cells and size

 - Cleaner API: solver doesn’t have to manipulate [][]rune directly

 - Makes it Go-style OOP-like, improves readability

***Step 6 — Printing solution***

 ```
for r := 0; r < b.Size; r++ {
    for c := 0; c < b.Size; c++ {
        fmt.Print(string(b.Cells[r][c]))
    }
    fmt.Println()
}
```

Simple row-by-row print

Uses letters assigned to tetrominoes for clarity


## Connectivity check (`isConnected`)

This function verifies that the 4 blocks of a tetromino form **one single connected shape** using only **side adjacency** (up, down, left, right).

Diagonal contact does **not** count.

---

### Function

```go
func isConnected(blocks []Block) bool {
	connections := 0

	for i := 0; i < len(blocks); i++ {
		for j := i + 1; j < len(blocks); j++ {

			dx := blocks[i].X - blocks[j].X
			if dx < 0 {
				dx = -dx
			}
			dy := blocks[i].Y - blocks[j].Y
			if dy < 0 {
				dy = -dy
			}

			// side-adjacent blocks
			if (dx == 1 && dy == 0) || (dx == 0 && dy == 1) {
				connections++
			}
		}
	}

	return connections >= 3
}
