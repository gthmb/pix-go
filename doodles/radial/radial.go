package main

import (
	"math"
	"math/rand"
	"os"

	"github.com/fogleman/gg"
)

const width = 1500
const height = 1500
const circleRadius = 500.0

func mapPoint(x1, x2, y1, y2, xt float64) float64 {
	return y1 + (xt-x1)*(y2-y1)/(x2-x1)
}

func drawCircleRing(dc *gg.Context, r, n float64) float64 {
	da := 2 * math.Pi / n
	individR := r * math.Sin(da/2)
	offset := n / 2
	for a := 0.0; a < 2*math.Pi; a += da {
		x := r * math.Cos(a+offset)
		y := r * math.Sin(a+offset)

		dc.SetRGBA255(rand.Intn(255), rand.Intn(255), rand.Intn(255), 30)
		dc.DrawRegularPolygon(8+rand.Intn(5), x, y, individR*2, rand.Float64())
		dc.FillPreserve()

		dc.SetRGBA255(rand.Intn(255), rand.Intn(255), rand.Intn(255), 10)
		dc.DrawRegularPolygon(8+rand.Intn(5), x, y, individR, rand.Float64())
		dc.FillPreserve()

		if individR > 2 {
			dc.SetRGBA255(rand.Intn(255), rand.Intn(255), rand.Intn(255), 10)
			dc.DrawRegularPolygon(8+rand.Intn(5), x, y, individR/10, rand.Float64())
		}

		dc.SetRGBA255(rand.Intn(255), rand.Intn(255), rand.Intn(255), 90)
		dc.Stroke()
	}
	return individR
}

func main() {
	dc := gg.NewContext(width, height)

	dc.SetRGB(245./255, 245./255, 245./255)
	dc.Clear()
	dc.SetRGB(0, 0, 0)
	dc.SetLineWidth(1)
	dc.Translate(width/2, height/2)

	r := 17.
	for i := 5.; i < 50; i++ {
		r += 2.1 * drawCircleRing(dc, r, 2*i)
	}

	dc.Stroke()
	dc.SavePNG(os.Args[1])
}
