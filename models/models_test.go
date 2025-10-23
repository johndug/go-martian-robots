package models

import "testing"

func TestPosition(t *testing.T) {
	pos := Position{X: 5, Y: 10}

	if pos.X != 5 {
		t.Errorf("Expected X=5, got %d", pos.X)
	}
	if pos.Y != 10 {
		t.Errorf("Expected Y=10, got %d", pos.Y)
	}
}

func TestGrid_IsInBounds(t *testing.T) {
	grid := Grid{Width: 5, Height: 3}

	// Valid positions
	validPositions := []Position{
		{0, 0}, {5, 3}, {2, 1}, {3, 2},
	}

	for _, pos := range validPositions {
		if !grid.IsInBounds(pos) {
			t.Errorf("Position (%d,%d) should be in bounds", pos.X, pos.Y)
		}
	}

	// Invalid positions
	invalidPositions := []Position{
		{-1, 0}, {0, -1}, {6, 3}, {5, 4},
	}

	for _, pos := range invalidPositions {
		if grid.IsInBounds(pos) {
			t.Errorf("Position (%d,%d) should be out of bounds", pos.X, pos.Y)
		}
	}
}
