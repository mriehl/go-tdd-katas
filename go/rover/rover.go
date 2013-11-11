package rover

import (
	"fmt"
)

const (
	North = iota
	East  = iota
	South = iota
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

// rover.New()
func New(coords Coordinates, facing int) *Rover {
	rover := new(Rover)
	rover.Coords = coords
	rover.Facing = facing
	return rover
}

func (rover *Rover) TurnRight() {
	switch direction := rover.Facing; direction {
	case North:
		rover.Facing = East
	case East:
		rover.Facing = South
	case South:
		rover.Facing = West
	case West:
		rover.Facing = North
	default:
		panic(fmt.Sprintf("WTF is this direction: %d", direction))
	}
}

func (rover *Rover) TurnLeft() {
	// turn left = turn right 3 times
	// screw efficiency, let's be DRY
	for i := 0; i < 3; i++ {
		rover.TurnRight()
	}
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
