package main

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"image"
	"image/color"
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

func TestShouldReturnSliceOfLengthDx(t *testing.T) {
	ySlice := Pic(2, 4)
	assert.Equal(t, len(ySlice), 2)
}

func TestShouldReturnDySlicesOfLengthDx(t *testing.T) {
	ySlice := Pic(2, 4)
	for index, xSlice := range ySlice {
		assert.Equal(t, len(xSlice), 4, fmt.Sprintf("xSlice at index %d does not have expected length dX", index))
	}
}

func TestShouldFillWithBluescaleValues(t *testing.T) {
	bluescales := Pic(2, 2)
	assert.Equal(t, bluescales,
		[][]uint8{
			[]uint8{BluescaleValue(0, 0), BluescaleValue(0, 1)},
			[]uint8{BluescaleValue(1, 0), BluescaleValue(1, 1)}})
}

func TestShouldExposeRGBAColorModel(t *testing.T) {
	bluescales := Pic(2, 2)
	img := GoImage{bluescales}

	assert.Equal(t, color.RGBAModel, img.ColorModel())
}

func TestShouldExposeBounds(t *testing.T) {
	bluescales := Pic(4, 6)
	img := GoImage{bluescales}

	assert.Equal(t,
		image.Rectangle{Min: image.Point{X: 0, Y: 0}, Max: image.Point{X: 4, Y: 6}},
		img.Bounds())
}

func TestShouldExposeColorValuesBasedOnMappingValue(t *testing.T) {
	bluescales := Pic(2, 2)
	img := GoImage{bluescales}

	for x := range bluescales {
		for y := range bluescales[x] {
			assert.Equal(t,
				color.RGBA{bluescales[x][y], bluescales[x][y], 255, 255},
				img.At(x, y))
		}
	}
}
