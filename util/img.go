package util

import "image/color"

func Mismatch(p1 color.Color, p2 color.Color) bool {
	r1, g1, b1, _ := p1.RGBA()
	r2, g2, b2, _ := p2.RGBA()
	diffR := absDiff(r1, r2)
	diffG := absDiff(g1, g2)
	diffB := absDiff(b1, b2)
	diffTot := diffR + diffG + diffB
	return diffTot > 20000
}

func absDiff(u1, u2 uint32) int {
	u := int(u1) - int(u2)
	if u > 0 {
		return u
	}
	return -u
}
