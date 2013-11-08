package main

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

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

/*
 *	Should use (x + y) / 2 as bluescale values
 */
func TestShouldFillWithInterestingValues(t *testing.T) {
	bluescales := Pic(4, 4)
	assert.Equal(t, bluescales,
		[][]uint8{[]uint8{0x0, 0x0, 0x1, 0x1}, []uint8{0x0, 0x1, 0x1, 0x2},
			[]uint8{0x1, 0x1, 0x2, 0x2}, []uint8{0x1, 0x2, 0x2, 0x3}})
}
