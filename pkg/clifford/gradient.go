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

type Gradient []GradientRec

func GradientHexSlice(tokens []string) (Gradient, error) {
	grad := make(Gradient, len(tokens))
	step := 1 / float64(len(tokens)-1)
	for i, tok := range tokens {
		c, err := colorful.Hex(tok)
		if err != nil {
			return nil, fmt.Errorf("invalid hex color: %w", err)
		}
		grad[i] = GradientRec{
			Col: c,
			Pos: float64(i) * step,
		}
	}
	return grad, nil
}

func GradientHexString(s string) (Gradient, error) {
	tokens := strings.Split(s, ",")
	if len(tokens) < 2 {
		return nil, fmt.Errorf("invalid gradient, requires at least 2 comma-separated, hex-formatted colors")
	}
	return GradientHexSlice(tokens)
}

func (g Gradient) Interp(t float64) colorful.Color {
	for i := 0; i < len(g)-1; i++ {
		c1 := g[i]
		c2 := g[i+1]
		if c1.Pos <= t && t <= c2.Pos {
			t := (t - c1.Pos) / (c2.Pos - c1.Pos)
			return c1.Col.BlendHcl(c2.Col, t).Clamped()
		}
	}
	return g[len(g)-1].Col
}
