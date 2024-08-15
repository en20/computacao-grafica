package transform

import (
	"cg-go/src/matrix"
	"cg-go/src/vec"
)

func TransformPoint(mtx [][]float64, point vec.Vec2D) vec.Vec2D {
	pointScaled := matrix.MatrixMult(mtx, point.ToTransposedXY1())
	return vec.NewVec2D(pointScaled[0][0], pointScaled[1][0])
}
