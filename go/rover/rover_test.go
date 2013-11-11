package rover

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewRoverShouldBeAtStartingPoint(t *testing.T) {
	rover := new(Rover)
	assert.Equal(t, rover.Coords, Coordinates{0, 0})
}

func TestRoverShouldBeFacingInitialDirection(t *testing.T) {
	var rover = Rover{Facing: South}
	assert.Equal(t, rover.Facing, South)
}

func TestShouldAdvanceNorth(t *testing.T) {
	var rover = &Rover{Facing: North}
	rover.Advance()
	assert.Equal(t, rover.Coords, Coordinates{0, 1})
}

func TestShouldAdvanceEast(t *testing.T) {
	var rover = &Rover{Facing: East}
	rover.Advance()
	assert.Equal(t, rover.Coords, Coordinates{1, 0})
}

func TestShouldAdvanceWest(t *testing.T) {
	var rover = &Rover{Facing: West}
	rover.Advance()
	assert.Equal(t, rover.Coords, Coordinates{-1, 0})
}

func TestShouldAdvanceSouth(t *testing.T) {
	var rover = &Rover{Facing: South}
	rover.Advance()
	assert.Equal(t, rover.Coords, Coordinates{0, -1})
}

func TestShouldRetreatNorth(t *testing.T) {
	var rover = &Rover{Facing: North}
	rover.Retreat()
	assert.Equal(t, rover.Coords, Coordinates{0, -1})
}

func TestShouldRetreatEast(t *testing.T) {
	var rover = &Rover{Facing: East}
	rover.Retreat()
	assert.Equal(t, rover.Coords, Coordinates{-1, 0})
}

func TestShouldRetreatWest(t *testing.T) {
	var rover = &Rover{Facing: West}
	rover.Retreat()
	assert.Equal(t, rover.Coords, Coordinates{1, 0})
}

func TestShouldRetreatSouth(t *testing.T) {
	var rover = &Rover{Facing: South}
	rover.Retreat()
	assert.Equal(t, rover.Coords, Coordinates{0, 1})
}
