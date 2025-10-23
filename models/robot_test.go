package models

import "testing"

func TestRobot_Command(t *testing.T) {
	grid := Grid{Width: 5, Height: 3}
	robot := Robot{
		Position: Position{X: 1, Y: 1},
		Facing:   "N",
		Grid:     grid,
		Lost:     false,
	}

	// Test turn left
	robot.Command("L")
	if robot.Facing != "W" {
		t.Errorf("Expected facing W after L, got %s", robot.Facing)
	}

	robot.Command("L")
	if robot.Facing != "S" {
		t.Errorf("Expected facing S after L, got %s", robot.Facing)
	}

	robot.Command("L")
	if robot.Facing != "E" {
		t.Errorf("Expected facing E after L, got %s", robot.Facing)
	}

	robot.Command("L")
	if robot.Facing != "N" {
		t.Errorf("Expected facing N after L, got %s", robot.Facing)
	}
}

func TestRobot_TurnRight(t *testing.T) {
	grid := Grid{Width: 5, Height: 3}
	robot := Robot{
		Position: Position{X: 1, Y: 1},
		Facing:   "N",
		Grid:     grid,
		Lost:     false,
	}

	robot.Command("R")
	if robot.Facing != "E" {
		t.Errorf("Expected facing E after R, got %s", robot.Facing)
	}

	robot.Command("R")
	if robot.Facing != "S" {
		t.Errorf("Expected facing S after R, got %s", robot.Facing)
	}

	robot.Command("R")
	if robot.Facing != "W" {
		t.Errorf("Expected facing W after R, got %s", robot.Facing)
	}

	robot.Command("R")
	if robot.Facing != "N" {
		t.Errorf("Expected facing N after R, got %s", robot.Facing)
	}
}

func TestRobot_Forward(t *testing.T) {
	grid := Grid{Width: 5, Height: 3}

	// Test North movement
	robot := Robot{
		Position: Position{X: 2, Y: 2},
		Facing:   "N",
		Grid:     grid,
		Lost:     false,
	}
	robot.Command("F")
	if robot.Position.X != 2 || robot.Position.Y != 3 {
		t.Errorf("Expected position (2,3), got (%d,%d)", robot.Position.X, robot.Position.Y)
	}

	// Test East movement
	robot.Facing = "E"
	robot.Command("F")
	if robot.Position.X != 3 || robot.Position.Y != 3 {
		t.Errorf("Expected position (3,3), got (%d,%d)", robot.Position.X, robot.Position.Y)
	}

	// Test South movement
	robot.Facing = "S"
	robot.Command("F")
	if robot.Position.X != 3 || robot.Position.Y != 2 {
		t.Errorf("Expected position (3,2), got (%d,%d)", robot.Position.X, robot.Position.Y)
	}

	// Test West movement
	robot.Facing = "W"
	robot.Command("F")
	if robot.Position.X != 2 || robot.Position.Y != 2 {
		t.Errorf("Expected position (2,2), got (%d,%d)", robot.Position.X, robot.Position.Y)
	}
}

func TestRobot_Lost(t *testing.T) {
	grid := Grid{Width: 5, Height: 3}

	// Test robot getting lost at boundary
	robot := Robot{
		Position: Position{X: 5, Y: 3},
		Facing:   "N",
		Grid:     grid,
		Lost:     false,
	}

	robot.Command("F")
	if !robot.Lost {
		t.Error("Robot should be lost after moving out of bounds")
	}

	expectedMsg := "5, 3, N LOST"
	if robot.CoordMsg() != expectedMsg {
		t.Errorf("Expected message '%s', got '%s'", expectedMsg, robot.CoordMsg())
	}
}

func TestRobot_ScentPrevention(t *testing.T) {
	grid := Grid{Width: 5, Height: 3}

	// Clear scents for clean test
	scents = make(map[string]bool)

	// First robot gets lost
	robot1 := Robot{
		Position: Position{X: 5, Y: 3},
		Facing:   "N",
		Grid:     grid,
		Lost:     false,
	}
	robot1.Command("F")
	if !robot1.Lost {
		t.Error("First robot should be lost")
	}

	// Second robot at same position/direction should NOT get lost due to scent
	robot2 := Robot{
		Position: Position{X: 5, Y: 3},
		Facing:   "N",
		Grid:     grid,
		Lost:     false,
	}
	robot2.Command("F")
	if robot2.Lost {
		t.Error("Second robot should NOT be lost due to scent")
	}

	expectedMsg := "5, 3, N"
	if robot2.CoordMsg() != expectedMsg {
		t.Errorf("Expected message '%s', got '%s'", expectedMsg, robot2.CoordMsg())
	}
}

func TestRobot_CoordMsg(t *testing.T) {
	grid := Grid{Width: 5, Height: 3}

	// Test normal position message
	robot := Robot{
		Position: Position{X: 2, Y: 3},
		Facing:   "S",
		Grid:     grid,
		Lost:     false,
	}

	expectedMsg := "2, 3, S"
	if robot.CoordMsg() != expectedMsg {
		t.Errorf("Expected message '%s', got '%s'", expectedMsg, robot.CoordMsg())
	}

	// Test lost position message
	robot.Lost = true
	robot.ScentPosition = Position{X: 1, Y: 1}
	robot.Facing = "E"

	expectedLostMsg := "1, 1, E LOST"
	if robot.CoordMsg() != expectedLostMsg {
		t.Errorf("Expected lost message '%s', got '%s'", expectedLostMsg, robot.CoordMsg())
	}
}
