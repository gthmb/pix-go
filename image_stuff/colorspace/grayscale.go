package colorspace

import (
	"image"
	"image/color"
)

// GreyScaleFilter implements image.Image and overrides the At method
type GreyScaleFilter struct {
	image.Image
}

// At is the pixel printer. converts a pixel to grayscale
func (f *GreyScaleFilter) At(x, y int) color.Color {
	r, g, b, a := f.Image.At(x, y).RGBA()

	// get the luminance: Y = 0.21R + 0.72G + 0.07B
	grey := uint16(float64(r)*0.21 + float64(g)*0.72 + float64(b)*0.07)

	return color.RGBA64{
		R: grey,
		G: grey,
		B: grey,
		A: uint16(a),
	}
}
