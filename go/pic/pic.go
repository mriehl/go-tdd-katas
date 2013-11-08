package main

import (
	"code.google.com/p/go-tour/pic"
)

func Pic(dx, dy int) (bluescale [][]uint8) {
	bluescale = make([][]uint8, dy)

	for dxSliceIndex := range bluescale {
		bluescale[dxSliceIndex] = make([]uint8, dx)

		for pixelIndex := range bluescale[dxSliceIndex] {
			bluescale[dxSliceIndex][pixelIndex] = ((uint8(dxSliceIndex) +
				uint8(pixelIndex)) / uint8(2))
		}

	}

	return
}

func main() {
	pic.Show(Pic)
}
