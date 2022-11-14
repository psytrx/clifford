package clifford

import (
	"image"
	"math"
)

func RenderHistogram(hist Histogram, size int, grad *Gradient) image.Image {
	img := image.NewRGBA(image.Rect(0, 0, size, size))
	for i := 0; i < size*size; i++ {
		hits := hist.Bins[i]

		f := float64(hits) / float64(hist.Limit)
		ff := math.Log(1 + f*(math.E-1))
		c := grad.Interp(ff)

		ix := i % size
		iy := (i / size)
		img.Set(ix, iy, c)
	}
	return img
}
