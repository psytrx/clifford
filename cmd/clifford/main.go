package main

import (
	"clifford/pkg/clifford"
	"clifford/pkg/huemint"
	"fmt"
	"image"
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

	log.Println("fetching random gradient...")
	grad, err := randomGradient()
	if err != nil {
		log.Fatalf("could not get random gradient: %s", err)
	}

	a, b, c, d := 1.7, 1.7, 0.6, 1.2
	att := clifford.NewAttractor(a, b, c, d)

	// stabilize
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

func randomGradient() (clifford.Gradient, error) {
	payload := huemint.Payload{
		NumColors:   6,
		Temperature: 1.3,
		NumResults:  1,
		Adjacency:   []string{"0", "35", "45", "55", "65", "75", "35", "0", "15", "25", "35", "45", "45", "15", "0", "15", "25", "35", "55", "25", "15", "0", "15", "25", "65", "35", "25", "15", "0", "15", "75", "45", "35", "25", "15", "0"},
		Palette:     []string{"-", "-", "-", "-", "-", "-"},
		Mode:        huemint.ModeTransformer,
	}

	res, err := huemint.Colors(payload)
	if err != nil {
		return nil, fmt.Errorf("could not get colors from Huemint: %w", err)
	}

	grad, err := clifford.GradientFromSlice(res.Results[0].Palette)
	if err != nil {
		return nil, fmt.Errorf("could not create gradient from slice: %w", err)
	}

	return grad, nil
}

func writeImage(filename string, img image.Image) error {
	f, err := os.Create(filename)
	if err != nil {
		return fmt.Errorf("could not create output file: %v", err)
	}
	defer f.Close()

	if err := jpeg.Encode(f, img, &jpeg.Options{Quality: 100}); err != nil {
		return fmt.Errorf("could not encode JPEG: %v", err)
	}

	return nil
}
