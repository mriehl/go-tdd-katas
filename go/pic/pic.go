package main

import (
	"code.google.com/p/go-tour/pic"
	"image"
	"image/color"
)

type GoImage struct {
	Bluescale [][]uint8
}

func (img GoImage) ColorModel() color.Model {
	return color.RGBAModel
}

func (img GoImage) Bounds() image.Rectangle {
	width := len(img.Bluescale)
	height := len(img.Bluescale[0])
	return image.Rect(0, 0, width, height)
}

func (img GoImage) At(x, y int) color.Color {
	bluescaleValue := img.Bluescale[x][y]
	return color.RGBA{bluescaleValue, bluescaleValue, 255, 255}
}

func BluescaleValue(x, y int) (value uint8) {
	_x := uint8(x)
	_y := uint8(y)
	value = (_x + _y) / uint8(2)
	return
}

func Pic(dx, dy int) (bluescale [][]uint8) {
	bluescale = make([][]uint8, dx)

	for dxSliceIndex := range bluescale {
		bluescale[dxSliceIndex] = make([]uint8, dy)

		for pixelIndex := range bluescale[dxSliceIndex] {
			pixelValue := BluescaleValue(dxSliceIndex, pixelIndex)
			bluescale[dxSliceIndex][pixelIndex] = pixelValue
		}

	}

	return
}

func main() {
	img := GoImage{Pic(800, 600)}
	pic.ShowImage(img)
}
