package world

import (
	"cg-go/src/bitmap"
	"cg-go/src/colors"
	"cg-go/src/geo"
	"cg-go/src/vec"

	"github.com/hajimehoshi/ebiten/v2"
)

var centerCircle = vec.NewVec2D(Width/2, Height/2)
var centerEllipse = vec.NewVec2D(Width/3, Height/3)

func BitMapUpdate(ctx *ebiten.Image) {
	Mem.Clear(colors.ColorBlack)

	DrawC(Mem, vec.NewVec2D(30, 50))
	DrawG(Mem, vec.NewVec2D(90, 50))

	geo.NewRect(280, 180, vec.NewVec2D(Width/2, Height/2)).DrawBounds(Mem)
	bitmap.FloodFill(Mem, vec.NewVec2D(285, 185), colors.ColorPurple, colors.ColorWhite)

	geo.NewRect(296, 196, vec.NewVec2D(Width/2, Height/2)).DrawBounds(Mem)
	bitmap.FloodFill(Mem, vec.NewVec2D(5, 5), colors.ColorSilver, colors.ColorWhite)

	bitmap.DrawFilledEllipse(Mem, centerEllipse, 50, 20, colors.ColorPink, colors.ColorWhite)
	bitmap.BresenhamDrawFilledCircle(Mem, centerCircle, 50, colors.ColorWhite, colors.ColorSilver)

	geo.NewTriangle(100, 100, vec.NewVec2D(Width/2+100, Height/3)).DrawBounds(Mem)
	bitmap.FloodFill(Mem, vec.NewVec2D(Width/2+100, Height/3), colors.ColorTeal, colors.ColorWhite)

	geo.NewRect(90, 30, vec.NewVec2D(Width/5, Height-100)).DrawBounds(Mem)
	bitmap.FloodFill(Mem, vec.NewVec2D(Width/5, Height-100), colors.ColorOrange, colors.ColorWhite)

	Mem.Draw(ctx)
}
