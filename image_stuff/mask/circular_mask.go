package mask

import (
	"image"
	"image/color"
)

// CircularMask struct
type CircularMask struct {
	Source image.Image
	Center image.Point
	Radius int
}

// ColorModel ...
func (c *CircularMask) ColorModel() color.Model {
	return c.Source.ColorModel()
}

// Bounds gets the rect of an image
func (c *CircularMask) Bounds() image.Rectangle {
	return image.Rect(
		c.Center.X-c.Radius,
		c.Center.Y-c.Radius,
		c.Center.X+c.Radius,
		c.Center.Y+c.Radius,
	)
}

// At finds the color at a given point if it is within the mask radius,
// otherwise return a transparent value
func (c *CircularMask) At(x, y int) color.Color {
	xx := float64(x - c.Center.X)
	yy := float64(y - c.Center.Y)
	rr := float64(c.Radius)

	if xx*xx+yy*yy < rr*rr {
		return c.Source.At(x, y)
	}

	return color.Alpha{0}
}
