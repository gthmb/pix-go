package util

// Pixel ...
type Pixel struct {
	R int
	G int
	B int
	A int
}

// RGBA converts a Pixel to it's RGBA values
func (p Pixel) RGBA() (r, g, b, a int) {
	return p.R, p.G, p.B, p.A
}

// RGBAToPixel assembles a Pixel from the given RGBA values
func RGBAToPixel(r, g, b, a uint32) Pixel {
	return Pixel{int(r / 257), int(g / 257), int(b / 257), int(a / 257)}
}
