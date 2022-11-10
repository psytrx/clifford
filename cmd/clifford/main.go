package main

import (
	"image"
	"image/color"
	"image/jpeg"
	"log"
	"math"
	"math/rand"
	"os"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())

	a := -2 + 4*rand.Float64()
	b := -2 + 4*rand.Float64()
	c := -2 + 4*rand.Float64()
	d := -2 + 4*rand.Float64()

	x, y := 0.0, 0.0

	img := image.NewRGBA(image.Rect(0, 0, 1024, 1024))
	for i := 0; i < 1e6; i++ {
		nx := math.Sin(a*y) + c*math.Cos(a*x)
		ny := math.Sin(b*x) + d*math.Cos(b*y)

		x, y = nx, ny

		if i > 32 {
			ix := int(1024/2 + x*512/3)
			iy := int(1024/2 + y*512/3)
			img.Set(ix, iy, color.White)
		}
	}

	f, err := os.Create("./output.jpg")
	if err != nil {
		log.Fatalf("could not create output file: %v", err)
	}
	defer f.Close()

	if err := jpeg.Encode(f, img, &jpeg.Options{Quality: 100}); err != nil {
		log.Fatalf("could not encode JPEG: %v", err)
	}
}
