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

	att := clifford.NewRandomAttractor(-2, 2)
	// stabilize
	for i := 0; i < 1e6; i++ {
		att.Advance()
	}

	img := image.NewRGBA(image.Rect(0, 0, 1024, 1024))
	for i := 0; i < 1e6; i++ {
		att.Advance()

		if i > 32 {
			ix := int(1024/2 + att.X*512/3)
			iy := int(1024/2 + att.Y*512/3)
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
