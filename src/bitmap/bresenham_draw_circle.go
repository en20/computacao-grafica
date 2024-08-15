package bitmap

import (
	"cg-go/src/memory"
	"cg-go/src/vec"
	"image/color"
	"math"
)

func BresenhamDrawCircle(mem memory.Memory, center vec.Vec2D, radius float64, c color.RGBA) {
	x0, y0 := int(center.X), int(center.Y)
	x, y := int(math.Round(radius)), 0
	err := 0

	for x >= y {
		mem.SetPixel(x0+x, y0+y, c)
		mem.SetPixel(x0+y, y0+x, c)
		mem.SetPixel(x0-x, y0+y, c)
		mem.SetPixel(x0-y, y0+x, c)
		mem.SetPixel(x0-x, y0-y, c)
		mem.SetPixel(x0-y, y0-x, c)
		mem.SetPixel(x0+x, y0-y, c)
		mem.SetPixel(x0+y, y0-x, c)

		y++
		err += 1 + 2*y
		if 2*(err-x)+1 > 0 {
			x--
			err += 1 - 2*x
		}
	}
}
