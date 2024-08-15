package window

import (
	"cg-go/src/geo"
	"cg-go/src/matrix"
	"cg-go/src/transform"
	"cg-go/src/vec"
)

type Window struct {
	pi, pf vec.Vec2D
}

func New(pi, pf vec.Vec2D) *Window {
	return &Window{pi, pf}
}

func (w *Window) MapPoints(s *geo.GeometricShape, vp *Viewport) {
	MapPointsToWindow(s, w, vp)
}

func (w *Window) Center() vec.Vec2D {
	return w.pi.Plus(w.pf).ScalarDiv(2)
}

func (w *Window) Pf() vec.Vec2D {
	return w.pf
}

func (w *Window) Pi() vec.Vec2D {
	return w.pi
}

func (w *Window) Zoom(ratio float64) {
	center := w.Center()
	mtx := transform.NewScaledTranslatedMatrix(center.X, center.Y, ratio, ratio)
	w.pi = transform.TransformPoint(mtx, w.pi)
	w.pf = transform.TransformPoint(mtx, w.pf)
}

func (w *Window) Translate(delta vec.Vec2D) {
	mtx := transform.NewTranslateMatrix(delta.X, delta.Y)
	w.pi = transform.TransformPoint(mtx, w.pi)
	w.pf = transform.TransformPoint(mtx, w.pf)
}

func MapPointToWindow(point vec.Vec2D, wi, wf vec.Vec2D, vp *Viewport) vec.Vec2D {

	a := (vp.pf.X - vp.pi.X) / (wf.X - wi.X)
	b := (vp.pf.Y - vp.pi.Y) / (wf.Y - wi.Y)

	mtx := [][]float64{
		{a, 0, vp.pi.X - a*wi.X},
		{0, b, vp.pi.Y - b*wi.Y},
		{0, 0, 1},
	}

	newPoint := matrix.MatrixMult(mtx, point.ToTransposedXY1())
	return vec.NewVec2D(newPoint[0][0], newPoint[1][0])
}

func MapPointsToWindow(s *geo.GeometricShape, w *Window, vp *Viewport) {
	for i, p := range s.Vertices {
		s.Vertices[i] = MapPointToWindow(p, w.pi, w.pf, vp)
	}
}

func (w *Window) IsInsideWindow(point vec.Vec2D) bool {
	return point.X >= w.pi.X && point.X <= w.pf.X && point.Y >= w.pi.Y && point.Y <= w.pf.Y
}
