package scan

import (
	"cg-go/src/colors"
	"cg-go/src/geo"
	"cg-go/src/memory"
	"cg-go/src/vec"
)

func ScanlineGradient(mem memory.Memory, s *geo.GeometricShape) {
	if len(s.Vertices) < 3 {
		return
	}

	ymin, ymax := vec.GetMinMaxY(s.Vertices)

	for y := ymin; y <= ymax; y++ {
		var i []ScanlinePointGradient

		pi := s.Vertices[0]

		for p := 1; p < len(s.Vertices); p++ {
			pf := s.Vertices[p]
			currentColor := s.ColorVertices[p]

			if point := Intersection(y, pi, pf); point.X >= 0 {
				i = append(i, NewScanlinePointGradient(point.X, point.T, currentColor))
			}

			pi, pf = pf, pi
		}

		pf, currentColor := s.Vertices[0], s.ColorVertices[0]

		if point := Intersection(y, pi, pf); point.X >= 0 {
			i = append(i, NewScanlinePointGradient(point.X, point.T, currentColor))
		}

		for pi := 0; pi < len(i); pi += 2 {

			if len(i) == 1 {
				continue
			}

			x1, t1, startColor := i[pi].X, i[pi].T, i[pi].Color
			x2, t2, endColor := i[pi+1].X, i[pi+1].T, i[pi+1].Color

			if x2 < x1 {
				x1, x2, t1, t2 = x2, x1, t2, t1
				startColor, endColor = endColor, startColor
			}

			for xk := x1; xk <= x2; xk++ {
				ratio := float64(xk-x1) / float64(x2-x1)
				smoothRatio := t1 + (t2-t1)*ratio

				c := colors.InterpolateColors(startColor, endColor, smoothRatio)
				mem.SetPixel(xk, y, c)
			}
		}

	}

}

func ScanlineGradient2(mem memory.Memory, s *geo.GeometricShape) {
	if len(s.Vertices) < 3 {
		return
	}

	ymin, ymax := vec.GetMinMaxY(s.Vertices)

	for y := ymin; y <= ymax; y++ {
		var i []ScanlinePointGradient

		pi := s.Vertices[0]
		pixelColor := s.ColorVertices[0]

		for p := 1; p <= len(s.Vertices); p++ {
			pf := s.Vertices[p%len(s.Vertices)]
			pfColor := s.ColorVertices[p%len(s.Vertices)]

			if point := Intersection(y, pi, pf); point.X >= 0 {
				i = append(i, NewScanlinePointGradient(point.X, point.T, pixelColor))
			}

			pi = pf
			pixelColor = pfColor
		}

		if len(i) < 2 {
			continue
		}

		for pi := 0; pi < len(i); pi++ {
			point1 := i[pi]
			point2 := i[(pi+1)%len(i)]

			if point1.X == point2.X {
				continue
			}

			startX, endX := point1.X, point2.X
			startT, endT := point1.T, point2.T
			startColor, endColor := point1.Color, point2.Color

			if startX > endX {
				startX, endX = endX, startX
				startT, endT = endT, startT
				startColor, endColor = endColor, startColor
			}

			for x := startX; x <= endX; x++ {
				ratio := float64(x-startX) / float64(endX-startX)
				smoothRatio := startT + (endT-startT)*ratio

				// Interpolação das cores
				interpolatedColor := colors.InterpolateColors(startColor, endColor, smoothRatio)

				mem.SetPixel(x, y, interpolatedColor)
			}
		}
	}
}
