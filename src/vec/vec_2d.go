package vec

import (
	"fmt"
	"math"
)

type Vec2D struct {
	X, Y float64
}

func NewVec2D(x, y float64) Vec2D {
	return Vec2D{truncate(x), truncate(y)}
}

func Zeros() Vec2D {
	return Vec2D{0, 0}
}

func Ones() Vec2D {
	return Vec2D{1, 1}
}

func (v Vec2D) String() string {
	return fmt.Sprintf("(%f, %f)", v.X, v.Y)
}

func (v Vec2D) ToXY1() []float64 {
	return []float64{v.X, v.Y, 1}
}

func (v Vec2D) ToTransposedXY1() [][]float64 {
	return [][]float64{{v.X}, {v.Y}, {1}}
}

func (v Vec2D) IsZero() bool {
	return v.X == 0 && v.Y == 0
}

func (v Vec2D) Plus(vec Vec2D) Vec2D {
	return NewVec2D(v.X+vec.X, v.Y+vec.Y)
}

func (v Vec2D) Sub(vec Vec2D) Vec2D {
	return NewVec2D(v.X-vec.X, v.Y-vec.Y)
}

func (v Vec2D) ScalarSum(value float64) Vec2D {
	return NewVec2D(v.X+value, v.Y+value)
}

func (v Vec2D) ScalarSub(value float64) Vec2D {
	return v.ScalarSum(-value)
}

func (v Vec2D) ScalarMult(value float64) Vec2D {
	return NewVec2D(v.X*value, v.Y*value)
}

func (v Vec2D) ScalarDiv(value float64) Vec2D {
	return v.ScalarMult(1 / value)
}

func (v Vec2D) Cross(vec Vec2D) float64 {
	return v.X*vec.Y - v.Y*vec.X
}

func Distance(p1, p2 Vec2D) float64 {
	dx := p2.X - p1.X
	dy := p2.Y - p1.Y
	return dx*dx + dy*dy
}

func GetMinMaxY(vertices []Vec2D) (ymin int, ymax int) {
	ymin = math.MaxUint32
	ymax = 0

	for _, p := range vertices {
		if p.Y < float64(ymin) {
			ymin = int(math.Round(p.Y))
		}

		if p.Y > float64(ymax) {
			ymax = int(math.Round(p.Y))
		}
	}

	return
}

func truncate(num float64) float64 {
	scale := math.Pow10(5)
	truncated := math.Round(num*scale) / scale
	return truncated
}
