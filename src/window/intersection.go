package window

import (
	"cg-go/src/vec"
	"math"
)

func Intersection(pi, pf, linePi, linePf vec.Vec2D) (vec.Vec2D, float64, float64) {
	// Cálculo dos vetores de direção dos segmentos
	dir1 := pf.Sub(pi)
	dir2 := linePf.Sub(linePi)

	// Cálculo do determinante
	det := dir1.Cross(dir2)

	if det == 0 {
		return vec.NewVec2D(-1, -1), 0, 0
	}

	// Verifica se os segmentos são paralelos ou coincidentes
	if math.Abs(det) < 1e-8 {
		return vec.NewVec2D(-1, -1), 0, 0
	}

	// Cálculo dos vetores entre os pontos iniciais dos segmentos
	startDiff := linePi.Sub(pi)

	// Cálculo dos parâmetros de interseção
	t := startDiff.Cross(dir2) / det
	u := startDiff.Cross(dir1) / det

	// Verifica se a interseção ocorre dentro dos segmentos
	if t >= 0 && t <= 1 && u >= 0 && u <= 1 {
		// Cálculo das coordenadas do ponto de interseção
		// intersectionX := pi.X + t*dir1.X
		// intersectionY := pi.Y + t*dir1.Y
		point := pi.Plus(dir1.ScalarMult(t))
		return point, t, u
	}

	// Não há interseção dentro dos segmentos
	return vec.NewVec2D(-1, -1), 0, 0
}
