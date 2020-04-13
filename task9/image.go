package task9

import (
	"image"
	"image/color"
)

type Image struct {
	Width, Height int
	Clr           uint8
}

func (i *Image) Bounds() image.Rectangle {
	return image.Rect(0, 0, i.Width, i.Height)
}

func (i *Image) ColorModel() color.Model {
	return color.RGBAModel
}

func (i *Image) At(x, y int) color.Color {
	return color.RGBA{
		R: i.Clr + uint8(x),
		G: i.Clr + uint8(y),
		B: 255,
		A: 255,
	}
}
