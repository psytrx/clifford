package clifford

import (
	"math"
	"math/rand"
)

type Attractor struct {
	A, B, C, D float64
	X, Y       float64
}

func NewAttractor(a, b, c, d float64) Attractor {
	return Attractor{
		A: a,
		B: b,
		C: c,
		D: d,
		X: 0,
		Y: 0,
	}
}

func NewRandomAttractor(min, max float64) Attractor {
	a := uniform(min, max)
	b := uniform(min, max)
	c := uniform(min, max)
	d := uniform(min, max)
	return NewAttractor(a, b, c, d)
}

func (att Attractor) Bounds() (float64, float64) {
	hw := 1 + math.Abs(att.C)
	hh := 1 + math.Abs(att.D)
	return 2 * hw, 2 * hh
}

func (att *Attractor) Advance() {
	nx := math.Sin(att.A*att.Y) + att.C*math.Cos(att.A*att.X)
	ny := math.Sin(att.B*att.X) + att.D*math.Cos(att.B*att.Y)
	att.X, att.Y = nx, ny
}

func uniform(min, max float64) float64 {
	return min + (max-min)*rand.Float64()
}
