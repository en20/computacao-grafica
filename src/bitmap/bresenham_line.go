package bitmap

import (
	"cg-go/src/memory"
	"cg-go/src/vec"
	"image/color"
	"math"
)

func BresenhamLine(mem memory.Memory, p0 vec.Vec2D, p1 vec.Vec2D, c color.RGBA) {
	x0, y0 := int(math.Round(p0.X)), int(math.Round(p0.Y))
	x1, y1 := int(p1.X), int(p1.Y)

	deltaX := math.Abs(float64(x1 - x0))
	deltaY := math.Abs(float64(y1 - y0))

	dirX := 1
	if x0 > x1 {
		dirX = -1
	}

	dirY := 1
	if y0 > y1 {
		dirY = -1
	}

	err := deltaX - deltaY

	for {
		mem.SetPixel(x0, y0, c)

		if x0 == x1 && y0 == y1 {
			break
		}

		err2 := err * 2

		if err2 > -deltaY {
			err -= deltaY
			x0 += dirX
		}

		if err2 < deltaX {
			err += deltaX
			y0 += dirY
		}
	}
}
