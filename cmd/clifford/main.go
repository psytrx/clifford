package main

import (
	"clifford/pkg/clifford"
	"log"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())

	size := 1024
	stabSteps := 1e4
	steps := int(1e8)

	log.Println("fetching random gradient...")
	grad, err := randomGradient()
	if err != nil {
		log.Fatalf("could not get random gradient: %s", err)
	}

	a, b, c, d := 1.7, 1.7, 0.6, 1.2
	att := clifford.NewAttractor(a, b, c, d)

	log.Println("stabilizing attractor...")
	for i := 0; i < int(stabSteps); i++ {
		att.Advance()
	}

	log.Println("building histogram...")
	hist := clifford.NewHistogram(size, att)
	for i := 0; i < steps; i++ {
		att.Advance()
		hist.Inc(att.X, att.Y)
	}

	log.Println("rendering histogram...")
	img := clifford.RenderHistogram(hist, size, grad)

	log.Println("writing output image...")
	if err := writeImage("./output.jpg", img); err != nil {
		log.Fatalf("could not write image: %s", err)
	}
}
