package scan

import (
	"cg-go/src/vec"
	"math"
)

func Intersection(y int, pi, pf vec.Vec2D) ScanlinePoint {

	// se segmento horizontal, não tem interseção
	if pi.Y == pf.Y {
		return NewScanlinePoint(-1, -1)
	}

	// troca para garantir ponto inicial em cima
	if pi.Y > pf.Y {
		pi, pf = pf, pi
	}

	// calcula t
	t := (float64(y) - (pi.Y)) / (pf.Y - pi.Y)

	// calcula x
	if t > 0 && t <= 1 {
		return NewScanlinePoint(int(math.Round((pi.X)+t*(pf.X-pi.X))), t)
	}

	// sem interseção
	return NewScanlinePoint(-1, -1)
}

func IntersectionForTexture(y int, pi, pf, texturePi, texturePf vec.Vec2D) ScanlinePointTexture {

	if pi.Y == pf.Y {
		return NewScanlinePointTexture(-1, 0, 0, 0)
	}

	if pi.Y > pf.Y {
		pi, pf = pf, pi
		texturePi, texturePf = texturePf, texturePi
	}

	t := (float64(y) - (pi.Y)) / (pf.Y - pi.Y)

	if t > 0 && t <= 1 {
		x := math.Round((pi.X) + t*((pf.X)-(pi.X)))
		tx := texturePi.X + t*(texturePf.X-texturePi.X)
		ty := texturePi.Y + t*(texturePf.Y-texturePi.Y)

		return NewScanlinePointTexture(int(x), y, tx, ty)
	}

	return NewScanlinePointTexture(-1, 0, 0, 0)
}
