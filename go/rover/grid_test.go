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
