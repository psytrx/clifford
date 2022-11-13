package main

import (
	"fmt"
	"image"
	"image/jpeg"
	"os"
)

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
