package transform

import (
	"cg-go/src/geo"
	"cg-go/src/matrix"
	"cg-go/src/vec"
)

func NewScaleMatrix(sx, sy float64) [][]float64 {
	return [][]float64{{sx, 0, 0}, {0, sy, 0}, {0, 0, 1}}
}

func NewScaledTranslatedMatrix(dx, dy, sx, sy float64) [][]float64 {
	translateMatForward := NewTranslateMatrix(dx, dy)
	translateMatBack := NewTranslateMatrix(-dx, -dy)
	rawScaleMat := NewScaleMatrix(sx, sy)
	return matrix.MatrixMult(matrix.MatrixMult(translateMatForward, rawScaleMat), translateMatBack)
}

func ScaleVertices(sx, sy float64, s *geo.GeometricShape) {
	mtx := NewScaledTranslatedMatrix(float64(s.Vertices[0].X), float64(s.Vertices[0].Y), sx, sy)

	var scaled []vec.Vec2D

	for _, point := range s.Vertices {
		scaled = append(scaled, TransformPoint(mtx, point))
	}

	s.Vertices = scaled
}
