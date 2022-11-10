package main

import (
	"math"
	"math/rand"
)

func main() {
	a := -2 + 4*rand.Float64()
	b := -2 + 4*rand.Float64()
	c := -2 + 4*rand.Float64()
	d := -2 + 4*rand.Float64()

	x, y := 0.0, 0.0

	for i := 0; i < 1e6; i++ {
		nx := math.Sin(a*y) + c*math.Cos(a*x)
		ny := math.Sin(b*x) + d*math.Cos(b*y)
		x, y = nx, ny
	}
}
