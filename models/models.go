package models

import "fmt"

type Grid struct {
	Width  int
	Height int
}

func (g *Grid) IsInBounds(p Position) bool {
	return p.X >= 0 && p.X <= g.Width && p.Y >= 0 && p.Y <= g.Height
}

type Position struct {
	X int
	Y int
}

type Robot struct {
	Position      Position
	Facing        string
	Grid          Grid
	Lost          bool
	ScentPosition Position
}

var scents = make(map[string]bool)

func (r *Robot) Command(command string) {
	if r.Lost {
		return
	}

	switch command {
	case "F":
		r.Forward()
	case "L":
		r.TurnLeft()
	case "R":
		r.TurnRight()
	}
}

func (r *Robot) CoordMsg() string {
	if r.Lost {
		return fmt.Sprintf("%d, %d, %s LOST", r.ScentPosition.X, r.ScentPosition.Y, r.Facing)
	}
	return fmt.Sprintf("%d, %d, %s", r.Position.X, r.Position.Y, r.Facing)
}

func (r *Robot) Forward() {
	nextPos := r.Position
	switch r.Facing {
	case "N":
		nextPos.Y++
	case "E":
		nextPos.X++
	case "S":
		nextPos.Y--
	case "W":
		nextPos.X--
	}

	if !r.Grid.IsInBounds(nextPos) {
		key := fmt.Sprintf("%d,%d,%s", r.Position.X, r.Position.Y, r.Facing)
		if !scents[key] {
			scents[key] = true
			r.Lost = true
			r.ScentPosition = r.Position
		}
		return
	}

	r.Position = nextPos
}

func (r *Robot) TurnLeft() {
	switch r.Facing {
	case "N":
		r.Facing = "W"
	case "E":
		r.Facing = "N"
	case "S":
		r.Facing = "E"
	case "W":
		r.Facing = "S"
	}
}

func (r *Robot) TurnRight() {
	switch r.Facing {
	case "N":
		r.Facing = "E"
	case "E":
		r.Facing = "S"
	case "S":
		r.Facing = "W"
	case "W":
		r.Facing = "N"
	}
}
