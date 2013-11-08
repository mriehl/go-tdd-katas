package main

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestShouldGenerateBluescaleValues(t *testing.T) {
	for i := 0; i < 10; i++ {
		for j := 0; j < 10; j++ {
			expected := uint8((i + j) / 2)
			actual := BluescaleValue(i, j)
			assert.Equal(t, actual, expected)
		}
	}
}

func TestShouldReturnSliceOfLengthDy(t *testing.T) {
	ySlice := Pic(2, 4)
	assert.Equal(t, len(ySlice), 4)
}

func TestShouldReturnDySlicesOfLengthDx(t *testing.T) {
	ySlice := Pic(2, 4)
	for index, xSlice := range ySlice {
		assert.Equal(t, len(xSlice), 2, fmt.Sprintf("xSlice at index %d does not have expected length dX", index))
	}
}

func TestShouldFillWithBluescaleValues(t *testing.T) {
	bluescales := Pic(2, 2)
	assert.Equal(t, bluescales,
		[][]uint8{
			[]uint8{BluescaleValue(0, 0), BluescaleValue(0, 1)},
			[]uint8{BluescaleValue(1, 0), BluescaleValue(1, 1)}})
}
