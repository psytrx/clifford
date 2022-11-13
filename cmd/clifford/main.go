package main

import (
	"clifford/pkg/clifford"
	"log"
	"math"
	"math/rand"
	"time"
)

const (
	DotsPerCm = 119 // 300 DPI
	Size      = 20 * DotsPerCm
	Steps     = 1e8
)

func main() {
	rand.Seed(time.Now().UnixNano())

	log.Println("fetching random gradient...")
	grad, err := randomGradient()
	if err != nil {
		log.Fatalf("could not get random gradient: %s", err)
	}

	a, b, c, d := 1.7, 1.7, 0.6, 1.2
	att := clifford.NewAttractor(a, b, c, d)

	log.Println("stabilizing attractor...")
	for i := 0; i < 128; i++ {
		att.Advance()
	}

	log.Println("building histogram...")
	hist := clifford.NewHistogram(Size, math.Phi/2, att)
	for i := 0; i < Steps; i++ {
		att.Advance()
		hist.Inc(att.X, att.Y)
	}

	log.Println("rendering histogram...")
	img := clifford.RenderHistogram(hist, Size, grad)

	log.Println("writing output image...")
	if err := writeImage("./output.jpg", img); err != nil {
		log.Fatalf("could not write image: %s", err)
	}
}
