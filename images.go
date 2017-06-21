package main

import (
	"image"
	"image/color"

	"golang.org/x/tour/pic"
)

type Image struct {
	w, h int
	c    uint8
}

func (m Image) Bounds() image.Rectangle {
	return image.Rect(0, 0, m.w, m.h)
}

func (m Image) ColorModel() color.Model {
	return color.RGBAModel
}

func (m Image) At(x, y int) color.Color {
	return color.RGBA{m.c + uint8(x*y+x*y), m.c + uint8(x*y+x*y), 255, 255}
}

func main() {
	m := Image{100, 100, 100}
	pic.ShowImage(m)
}
