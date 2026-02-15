package tetromino

import "testing"

func TestBuildNewTetrominoValidSquare(t *testing.T) {
	lines := []string{
		"....",
		".##.",
		".##.",
		"....",
	}

	tet, err := BuildNewTetromino(lines, 'A')
	if err != nil {
		t.Fatalf("expected valid tetromino, got error: %v", err)
	}

	// after normalization, top-left must be (0,0)
	foundOrigin := false
	for _, b := range tet.Blocks {
		if b.X == 0 && b.Y == 0 {
			foundOrigin = true
		}
	}

	if !foundOrigin {
		t.Fatalf("normalized tetromino does not start at (0,0)")
	}
}

func TestBuildNewTetrominoInvalidBlockCount(t *testing.T) {
	lines := []string{
		"....",
		".##.",
		"....",
		"....",
	}

	_, err := BuildNewTetromino(lines, 'A')
	if err == nil {
		t.Fatalf("expected error for invalid number of blocks")
	}
}

func TestBuildNewTetrominoNotConnected(t *testing.T) {
	lines := []string{
		"#...",
		"...#",
		"...#",
		"....",
	}

	_, err := BuildNewTetromino(lines, 'A')
	if err == nil {
		t.Fatalf("expected error for non connected blocks")
	}
}
