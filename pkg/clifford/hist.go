package clifford

import "math"

type Histogram struct {
	n      int
	bounds Bounds
	scale  float64
	Bins   []int
	Limit  int
}

func NewHistogram(n int, bounds Bounds) Histogram {
	scale := float64(n) / math.Max(bounds.Width, bounds.Height)
	return Histogram{
		n:      n,
		bounds: bounds,
		scale:  scale,
		Bins:   make([]int, n*n),
		Limit:  0,
	}
}

func (hist *Histogram) Inc(x, y float64) {
	ix := int(float64(hist.n)/2 + x*hist.scale)
	iy := int(float64(hist.n/2) + y*hist.scale)
	idx := ix + iy*hist.n
	hist.Bins[idx]++
	hist.Limit = max(hist.Limit, hist.Bins[idx])
}

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}
