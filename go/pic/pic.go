package main

import (
	"code.google.com/p/go-tour/pic"
)

func bluescaleValue(x, y int) (value uint8) {
	_x := uint8(x)
	_y := uint8(y)
	value = (_x + _y) / uint8(2)
	return
}

func Pic(dx, dy int) (bluescale [][]uint8) {
	bluescale = make([][]uint8, dy)

	for dxSliceIndex := range bluescale {
		bluescale[dxSliceIndex] = make([]uint8, dx)

		for pixelIndex := range bluescale[dxSliceIndex] {
			bluescale[dxSliceIndex][pixelIndex] = bluescaleValue(dxSliceIndex, pixelIndex)
		}

	}

	return
}

func main() {
	pic.Show(Pic)
}
