package clifford

import "math"

type Attractor struct {
	A, B, C, D float64
	X, Y       float64
}

func NewAttractor(a, b, c, d float64) Attractor {
	att := Attractor{
		A: a,
		B: b,
		C: c,
		D: d,
		X: 0,
		Y: 0,
	}
	return att
}

func (att *Attractor) Advance() {
	nx := math.Sin(att.A*att.Y) + att.C*math.Cos(att.A*att.X)
	ny := math.Sin(att.B*att.X) + att.D*math.Cos(att.B*att.Y)
	att.X, att.Y = nx, ny
}
