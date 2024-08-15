package memory

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
)

type Memory interface {
	SetPixel(x, y int, color color.RGBA)
	GetPixel(x, y int) color.RGBA
	Draw(ctx *ebiten.Image)
	Clear(color color.RGBA)
	Width() int
	Height() int
}
