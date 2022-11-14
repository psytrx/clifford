package clifford

import (
	"fmt"
	"strings"

	"github.com/lucasb-eyer/go-colorful"
)

type GradientRec struct {
	Col colorful.Color
	Pos float64
}

type Gradient struct {
	table []GradientRec
	cache map[float64]colorful.Color
}

func GradientHexSlice(tokens []string) (*Gradient, error) {
	table := make([]GradientRec, len(tokens))
	step := 1 / float64(len(tokens)-1)
	for i, tok := range tokens {
		c, err := colorful.Hex(tok)
		if err != nil {
			return nil, fmt.Errorf("invalid hex color: %w", err)
		}
		table[i] = GradientRec{
			Col: c,
			Pos: float64(i) * step,
		}
	}
	grad := Gradient{
		table: table,
		cache: make(map[float64]colorful.Color),
	}
	return &grad, nil
}

func GradientHexString(s string) (*Gradient, error) {
	tokens := strings.Split(s, ",")
	if len(tokens) < 2 {
		return nil, fmt.Errorf("invalid gradient, requires at least 2 comma-separated, hex-formatted colors")
	}
	return GradientHexSlice(tokens)
}

func (g Gradient) Interp(t float64) colorful.Color {
	for i := 0; i < len(g.table)-1; i++ {
		c1 := g.table[i]
		c2 := g.table[i+1]
		if c1.Pos <= t && t <= c2.Pos {
			t := (t - c1.Pos) / (c2.Pos - c1.Pos)
			return c1.Col.BlendLuv(c2.Col, t).Clamped()
		}
	}
	return g.table[len(g.table)-1].Col
}
