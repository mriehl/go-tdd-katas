package rover

import (
	"fmt"
)

const (
	North = iota
	South = iota
	East  = iota
	West  = iota
)

func Opposite(direction int) (newDirection int) {
	switch direction {
	case North:
		return South
	case South:
		return North
	case East:
		return West
	case West:
		return East
	default:
		panic(fmt.Sprintf("WTF is this direction: %d", direction))
	}
}

type Coordinates struct {
	X int
	Y int
}

type Rover struct {
	Coords Coordinates
	Facing int
}

func (rover *Rover) Advance() {
	switch direction := rover.Facing; direction {
	case North:
		rover.Coords.Y += 1
	case East:
		rover.Coords.X += 1
	case West:
		rover.Coords.X -= 1
	case South:
		rover.Coords.Y -= 1
	}
}

func (rover *Rover) Retreat() {
	facing := rover.Facing
	rover.Facing = Opposite(rover.Facing)
	rover.Advance()
	rover.Facing = facing
}
