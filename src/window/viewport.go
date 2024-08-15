package window

import (
	"cg-go/src/bitmap"
	"cg-go/src/memory"
	"cg-go/src/vec"
	"image/color"
)

type Viewport struct {
	pi, pf vec.Vec2D
}

func NewViewport(pi, pf vec.Vec2D) *Viewport {
	return &Viewport{pi, pf}
}

func (vp *Viewport) Center() vec.Vec2D {
	return vp.pi.Plus(vp.pf).ScalarDiv(2)
}

func (vp *Viewport) DrawBounds(mem memory.Memory) {
	bitmap.BresenhamLine(mem, vp.pi, vec.NewVec2D(vp.pf.X, vp.pi.Y), color.RGBA{255, 255, 255, 255})
	bitmap.BresenhamLine(mem, vec.NewVec2D(vp.pf.X, vp.pi.Y), vp.pf, color.RGBA{255, 255, 255, 255})
	bitmap.BresenhamLine(mem, vp.pf, vec.NewVec2D(vp.pi.X, vp.pf.Y), color.RGBA{255, 255, 255, 255})
	bitmap.BresenhamLine(mem, vec.NewVec2D(vp.pi.X, vp.pf.Y), vp.pi, color.RGBA{255, 255, 255, 255})
}
