package clifford

func StableAttractor(min, max float64, threshold uint) Attractor {
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

		if threshold*hist.Limit > uint(n) {
			continue
		}
		return att
	}
}
