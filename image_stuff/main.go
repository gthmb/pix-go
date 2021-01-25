package main

import (
	"image"
	_ "image/jpeg" // Register JPEG format
	"image/png"
	"os"

	"github.com/gthmb/pix-go/image_stuff/colorspace"
	"github.com/gthmb/pix-go/image_stuff/mask"
)

func min(a int, b int) int {
	if a > b {
		return b
	}
	return a
}

func main() {
	infile, err := os.Open(os.Args[1])

	if err != nil {
		panic(err.Error)
	}

	defer infile.Close()

	src, _, err := image.Decode(infile)

	if err != nil {
		panic(err.Error)
	}

	bounds := src.Bounds()

	dest := &colorspace.GreyScaleFilter{Image: src}

	mask := &mask.CircularMask{
		Source: dest,
		Center: image.Point{X: bounds.Max.X / 2, Y: bounds.Max.Y / 2},
		Radius: min(bounds.Max.X, bounds.Max.Y) / 2,
	}

	if err != nil {
		panic(err.Error)
	}

	outfile, err := os.Create(os.Args[2])
	defer outfile.Close()
	png.Encode(outfile, mask)
}
