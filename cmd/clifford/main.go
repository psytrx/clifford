package main

import (
	"clifford/pkg/clifford"
	"image"
	"image/color"
	"image/jpeg"
	"log"
	"math/rand"
	"os"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())

	size := 512
	steps := int(1e7)

	a, b, c, d := 1.7, 1.7, 0.6, 1.2
	att := clifford.NewAttractor(a, b, c, d)
	// att := clifford.NewRandomAttractor(-2, 2)
	log.Println(att)

	// stabilize
	for i := 0; i < steps; i++ {
		att.Advance()
	}

	hist := clifford.NewHistogram(size, att.Bounds())

	for i := 0; i < steps; i++ {
		att.Advance()
		hist.Inc(att.X, att.Y)
	}

	img := image.NewRGBA(image.Rect(0, 0, size, size))
	for i := 0; i < size*size; i++ {
		hits := hist.Bins[i]
		if hits < 1 {
			continue
		}

		f := float64(hits) / float64(hist.Limit)
		int := uint8(255 * f)
		c := color.RGBA{int, int, int, 255}

		ix := i % size
		iy := (i / size)
		img.Set(ix, iy, c)
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
