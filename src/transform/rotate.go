package transform

import (
	"cg-go/src/geo"
	"cg-go/src/matrix"
	"cg-go/src/vec"
	"math"
)

func NewRotateMatrix(angle float64) [][]float64 {
	rad := angle * math.Pi / 180

	return [][]float64{
		{math.Cos(rad), -math.Sin(rad), 0},
		{math.Sin(rad), math.Cos(rad), 0},
		{0, 0, 1},
	}
}

func RotateVertices(angle float64, s *geo.GeometricShape) {
	mtx := NewRotateMatrix(angle)

	var rotated []vec.Vec2D

	for _, point := range s.Vertices {
		rotated = append(rotated, TransformPoint(mtx, point))
	}

	s.Vertices = rotated
}

func NewRotateMatriWithPivot(angle float64, pivot vec.Vec2D) [][]float64 {
	rad := angle * math.Pi / 180
	cos := math.Cos(rad)
	sin := math.Sin(rad)

	return [][]float64{
		{cos, -sin, pivot.X - pivot.X*cos + pivot.Y*sin},
		{sin, cos, pivot.Y - pivot.X*sin - pivot.Y*cos},
		{0, 0, 1},
	}
}

func RotateVerticesOnPivot(angle float64, pivot vec.Vec2D, s *geo.GeometricShape) {
	translate1 := NewTranslateMatrix(-pivot.X, -pivot.Y)

	rotate := NewRotateMatrix(angle)

	translate2 := NewTranslateMatrix(pivot.X, pivot.Y)

	mtx := matrix.MatrixMult(translate2, matrix.MatrixMult(rotate, translate1))

	// Aplica a transformação a cada ponto do polígono
	var rotated []vec.Vec2D

	for _, point := range s.Vertices {
		rotated = append(rotated, TransformPoint(mtx, point))
	}

	s.Vertices = rotated
}
