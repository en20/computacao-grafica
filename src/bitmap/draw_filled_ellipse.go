package bitmap

import (
	"cg-go/src/memory"
	"cg-go/src/vec"
	"image/color"
)

func DrawFilledEllipse(mem memory.Memory, center vec.Vec2D, a, b float64, bgColor, borderColor color.RGBA) {
	DrawEllipse(mem, center, a, b, borderColor)
	FloodFill(mem, center, bgColor, borderColor)
}
