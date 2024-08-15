package world

import (
	"cg-go/src/colors"
	"cg-go/src/transform"
	"cg-go/src/vec"

	"github.com/hajimehoshi/ebiten/v2"
)

func ShapesUpdate(ctx *ebiten.Image) {

	Mem.Clear(colors.ColorBlack)

	DrawStars(Mem)
	DrawMeteors(Mem)
	MapObjectsToVP(MainViewport)

	

	Mem.Draw(ctx)

	if ebiten.IsKeyPressed(ebiten.KeyZ) {
		Win.Zoom(0.95)
	}

	if ebiten.IsKeyPressed(ebiten.KeyX) {
		Win.Zoom(1.05)
	}
	RotateMeteors()
	transform.RotateVerticesOnPivot(10, blackHole.Center(), blackHole)
	transform.RotateVerticesOnPivot(0.1, terran.Center(), terran)
	transform.RotateVerticesOnPivot(-0.5, lava.Center(), lava)
	transform.RotateVerticesOnPivot(0.5, ice.Center(), ice)
	transform.RotateVerticesOnPivot(2, lava.Center(), baren)
	transform.TranslateVertices(vec.NewVec2D(10, 5), shootingStar)
	transform.TranslateVertices(vec.NewVec2D(0, 1), rect)
}
