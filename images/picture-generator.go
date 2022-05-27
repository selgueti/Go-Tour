package main

import (
	"image"
	"image/color"

	"golang.org/x/tour/pic"
)

type Image struct {
	w, h int
}

func (pic Image) ColorModel() color.Model {
	return color.RGBAModel
}

func (pic Image) Bounds() image.Rectangle {
	return image.Rect(0, 0, pic.w, pic.h)
}

func (pic Image) At(x, y int) color.Color {
	var c uint8 = uint8(x) ^ uint8(y)
	return color.RGBA{c, c, 255, 255}
}

func main() {
	m := Image{256, 256}
	pic.ShowImage(m)
}
