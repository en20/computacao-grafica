package scan

import (
	"cg-go/src/geo"
	"cg-go/src/memory"
	"cg-go/src/tex"
	"cg-go/src/vec"

	"image/color"
)

func ScanlineTexture(mem memory.Memory, s *geo.GeometricShape, texture [][]color.RGBA) {
	if len(s.Vertices) < 3 {
		return
	}

	if len(s.Vertices) != len(s.TextureVertices) {
		return
	}

	ymin, ymax := vec.GetMinMaxY(s.Vertices)

	for y := ymin; y <= ymax; y++ {
		var intersections []ScanlinePointTexture

		pi := s.Vertices[0]
		tPi := s.TextureVertices[0]

		for p := 1; p < len(s.Vertices); p++ {
			pf := s.Vertices[p]
			tPf := s.TextureVertices[p]

			if point := IntersectionForTexture(y, pi, pf, tPi, tPf); point.X >= 0 {
				intersections = append(intersections, point)
			}

			pi = pf
			tPi = tPf
		}

		pf := s.Vertices[0]
		tPf := s.TextureVertices[0]

		if point := IntersectionForTexture(y, pi, pf, tPi, tPf); point.X >= 0 {
			intersections = append(intersections, point)
		}

		for pi := 0; pi < len(intersections); pi += 2 {

			if len(intersections) == 1 {
				continue
			}

			p1, p2 := intersections[pi], intersections[pi+1]

			if p2.X < p1.X {
				p1, p2 = p2, p1
			}

			for xk := p1.X; xk <= p2.X; xk++ {
				var pc float64

				if p2.X != p1.X {
					pc = float64(xk-p1.X) / float64(p2.X-p1.X)
				}

				tx := p1.Tx + pc*(p2.Tx-p1.Tx)
				ty := p1.Ty + pc*(p2.Ty-p1.Ty)
				color := tex.GetPixelFromTexture(texture, tx, ty)
				mem.SetPixel(xk, y, color)
			}
		}

	}

}
