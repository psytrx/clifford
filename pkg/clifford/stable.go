package clifford

func FindStableAttractor(min, max float64) Attractor {
	for {
		att := NewRandomAttractor(min, max)

		for i := 0; i < 128; i++ {
			att.Advance()
		}

		n := int(1e5)
		hist := NewHistogram(64, 1, att)
		for i := 0; i < n; i++ {
			att.Advance()
			hist.Inc(att.X, att.Y)
		}

		if 10*hist.Limit > uint(n) {
			continue
		}
		return att
	}
}
