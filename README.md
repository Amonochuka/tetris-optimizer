PASTABLE BLOCK OF README
# tetris-optimizer

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

#...

#...

#...

#...


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

```Run:
go run . sample.txt
```

Output:

ABBB.

ACCC.

A....

A....

(letters may vary depending on the input order)

## Project structure
.
├── main.go

├── board/

│   └── board.go

├── solver/

│   └── solver.go

└── tetromino/

    └── tetromino.go


## Tests

Basic unit tests are provided for tetromino validation.

Run all tests with:

go test ./...

## Notes

  - The program uses a backtracking solver.

  - Tetrominoes are normalized so their coordinates always start at (0,0).

  - Only the Go standard library is used.

