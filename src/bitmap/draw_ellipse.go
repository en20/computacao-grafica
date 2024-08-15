package bitmap

import (
	"cg-go/src/memory"
	"cg-go/src/vec"
	"image/color"
	"math"
)

func DrawEllipse(mem memory.Memory, center vec.Vec2D, a, b float64, borderColor color.RGBA) {
	x0 := int(math.Round(center.X))
	y0 := int(math.Round(center.Y))

	a2 := a * a
	b2 := b * b

	for x := -int(a); x <= int(a); x++ {
		// Equação da elipise para y
		y := int(math.Sqrt((1 - (float64(x*x) / a2)) * b2))
		mem.SetPixel(x0+x, y0+y, borderColor)
		mem.SetPixel(x0+x, y0-y, borderColor)
	}

	for y := -int(b); y <= int(b); y++ {
		// Equação da elipise para x
		x := int(math.Sqrt((1 - (float64(y*y) / b2)) * a2))
		mem.SetPixel(x0+x, y0+y, borderColor)
		mem.SetPixel(x0-x, y0+y, borderColor)
	}
}
