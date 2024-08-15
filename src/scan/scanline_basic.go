package scan

import (
	"cg-go/src/geo"
	"cg-go/src/memory"
	"cg-go/src/vec"
	"image/color"
)

func ScanlineBasic(mem memory.Memory, s *geo.GeometricShape, color color.RGBA) {
	if len(s.Vertices) < 3 {
		return
	}

	ymin, ymax := vec.GetMinMaxY(s.Vertices)

	for y := ymin; y <= ymax; y++ {
		var i []int

		pi := s.Vertices[0]

		for p := 1; p < len(s.Vertices); p++ {
			pf := s.Vertices[p]

			if point := Intersection(y, pi, pf); point.X >= 0 {
				i = append(i, point.X)
			}

			pi, pf = pf, pi
		}

		pf := s.Vertices[0]

		if point := Intersection(y, pi, pf); point.X >= 0 {
			i = append(i, point.X)
		}

		for pi := 0; pi < len(i); pi += 2 {
			if len(i) == 1 {
				continue
			}

			x1, x2 := i[pi], i[pi+1]

			if x2 < x1 {
				x1, x2 = x2, x1
			}

			for xk := x1; xk <= x2; xk++ {
				mem.SetPixel(xk, y, color)
			}
		}

	}

}
