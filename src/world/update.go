package world

import (
	"time"

	"github.com/hajimehoshi/ebiten/v2"
)

var start = time.Now()

func Update(ctx *ebiten.Image) {
	if start.Add(time.Second*5).Compare(time.Now()) <= 0 {
		ShapesUpdate(ctx)
	} else {
		BitMapUpdate(ctx)
	}

}
