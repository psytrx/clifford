package main

import (
	"clifford/pkg/clifford"
	"image/jpeg"
	"log"
	"math/rand"
	"os"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())

	size := 1024
	stabSteps := 1e4
	steps := int(1e7)

	grad, err := clifford.ParseGradient("#000000,#ff0000,#00ff00,#0000ff,#ffffff")
	if err != nil {
		log.Fatalf("could not parse gradient: %s", err)
	}

	a, b, c, d := 1.7, 1.7, 0.6, 1.2
	att := clifford.NewAttractor(a, b, c, d)

	// stabilize
	for i := 0; i < int(stabSteps); i++ {
		att.Advance()
	}

	hist := clifford.NewHistogram(size, att.Bounds())
	for i := 0; i < steps; i++ {
		att.Advance()
		hist.Inc(att.X, att.Y)
	}

	img := clifford.RenderHistogram(hist, size, grad)

	f, err := os.Create("./output.jpg")
	if err != nil {
		log.Fatalf("could not create output file: %v", err)
	}
	defer f.Close()

	if err := jpeg.Encode(f, img, &jpeg.Options{Quality: 100}); err != nil {
		log.Fatalf("could not encode JPEG: %v", err)
	}
}
