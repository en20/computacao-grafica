package bitmap

import (
	"cg-go/src/memory"
	"cg-go/src/vec"
	"image/color"
)

func BresenhamDrawFilledCircle(mem memory.Memory, center vec.Vec2D, radius float64, colorBorder, colorBg color.RGBA) {
	BresenhamDrawCircle(mem, center, radius, colorBorder)
	FloodFill(mem, center, colorBg, colorBorder)
}
