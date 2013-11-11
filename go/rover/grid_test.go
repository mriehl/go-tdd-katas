package rover

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestShouldCreateNewGrid(t *testing.T) {
	assert.Equal(t,
		NewGrid(4, 5).Field,
		[][]int{
			[]int{0, 0, 0, 0, 0},
			[]int{0, 0, 0, 0, 0},
			[]int{0, 0, 0, 0, 0},
			[]int{0, 0, 0, 0, 0}})
}

func TestShouldSnapRoverToGrid(t *testing.T) {
	grid := NewGrid(4, 5)
	rover := New(Coordinates{1, 2}, North)

	grid.Snap(rover)
	assert.Equal(t, grid.At(rover.Coords), ROVER)
}

func TestShouldOverflowWidth(t *testing.T) {
	grid := NewGrid(4, 4)
	newPosition := grid.OverflowPosition(Coordinates{3, 3}, 1, 0)
	assert.Equal(t, newPosition, Coordinates{0, 3})
}

func TestShouldOverflowHeight(t *testing.T) {
	grid := NewGrid(4, 4)
	newPosition := grid.OverflowPosition(Coordinates{3, 3}, 0, 1)
	assert.Equal(t, newPosition, Coordinates{3, 0})
}
