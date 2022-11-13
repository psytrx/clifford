package clifford

import "math"

type Histogram struct {
	n     int
	scale float64
	Bins  []uint
	Limit uint
}

func NewHistogram(n int, scale float64, att Attractor) Histogram {
	width, height := att.Bounds()
	s := scale * float64(n) / math.Max(width, height)
	return Histogram{
		n:     n,
		scale: s,
		Bins:  make([]uint, n*n),
		Limit: 0,
	}
}

func (hist *Histogram) Inc(x, y float64) {
	ix := int(float64(hist.n)/2 + x*hist.scale)
	iy := int(float64(hist.n/2) + y*hist.scale)
	idx := ix + iy*hist.n
	hist.Bins[idx]++
	hist.Limit = max(hist.Limit, hist.Bins[idx])
}

func max(a, b uint) uint {
	if a >= b {
		return a
	}
	return b
}
