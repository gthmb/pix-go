package main

import (
	"fmt"
	"image"
	"image/color"
	"math/rand"
	"os"
	"time"

	"github.com/fogleman/gg"
	"github.com/gthmb/pix-go/generative_art/util"
)

// copied from https://leanpub.com/generative-art-in-golang, and the slightly tweaked

var (
	strokeSize      = 2000.0
	alpha           = 0.1
	strokeReduction = 0.998
	alphaIncrease   = 0.06
)

func main() {
	file, err := os.Open(os.Args[1])
	defer file.Close()

	if err != nil {
		panic(err.Error)
	}

	img, format, err := image.Decode(file)

	if err != nil {
		panic(err.Error)
	}

	fmt.Printf("source image format: %s", format)

	bounds := img.Bounds()
	width, height := bounds.Max.X, bounds.Max.Y

	dc := gg.NewContext(width, height)
	dc.SetColor(color.Black)
	dc.DrawRectangle(0, 0, float64(width), float64(height))
	dc.FillPreserve()
	dc.SetLineWidth(0.0)

	rand.Seed(time.Now().Unix())

	for i := 0; i < 5000; i++ {
		rndX := rand.Float64() * float64(width)
		rndY := rand.Float64() * float64(height)
		pix := util.RGBAToPixel(img.At(int(rndX), int(rndY)).RGBA())

		edges := 3 + rand.Intn(5)

		dc.DrawRegularPolygon(edges, rndX, rndY, strokeSize, rand.Float64())

		dc.SetRGBA255(pix.R, pix.G, pix.B, int(alpha))
		dc.FillPreserve()
		dc.Stroke()

		strokeSize *= strokeReduction
		alpha += alphaIncrease
	}

	dc.SavePNG(os.Args[2])
}
