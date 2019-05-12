package parser

func areSameColours(colours1 colours, colours2 colours) bool {
	if len(colours1) != len(colours2) {
		return false
	}

	for i, colour1 := range colours1 {
		colour2 := colours2[i]
		if colour1.ColStart != colour2.ColStart ||
			colour1.ColEnd != colour2.ColEnd ||
			colour1.Lnum != colour2.Lnum ||
			colour1.Hex != colour2.Hex {
			return false
		}
	}

	return true
}
