package rover

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestShouldAcceptEmptyCommands(t *testing.T) {
	assert.Equal(t, ValidateCommands([]string{}), true)
}

func TestShouldAcceptValidCommands(t *testing.T) {
	assert.Equal(t, ValidateCommands([]string{"F", "B", "L", "R"}), true)
}

func TestShouldNotAcceptInvalidCommands(t *testing.T) {
	assert.Equal(t, ValidateCommands([]string{"F", "B", "X", "L", "R"}), false)
}

func TestRemoteShouldNotAcceptInvalidCommands(t *testing.T) {
	rover := New(Coordinates{0, 0}, East)
	remote := NewRemoteControl(rover)
	assert.Equal(t, remote.Send("X"), false)
}

func TestRemoteShouldAcceptValidCommands(t *testing.T) {
	rover := New(Coordinates{0, 0}, East)
	remote := NewRemoteControl(rover)
	assert.Equal(t, remote.Send("F", "B"), true)
}

func TestShouldDispatchCommandToRover(t *testing.T) {
	rover := New(Coordinates{0, 0}, East)
	remote := NewRemoteControl(rover)
	remote.Send("F")
	assert.Equal(t, rover.Coords, Coordinates{1, 0})
}

func TestShouldDispatchCommandsToRover(t *testing.T) {
	rover := New(Coordinates{0, 0}, North)
	remote := NewRemoteControl(rover)
	remote.Send("F", "F", "R")
	assert.Equal(t, rover.Coords, Coordinates{0, 2})
	assert.Equal(t, rover.Facing, East)

	remote.Send("L", "B")
	assert.Equal(t, rover.Facing, North)
	assert.Equal(t, rover.Coords, Coordinates{0, 1})
}
