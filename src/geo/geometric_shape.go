package geo

import (
	"cg-go/src/bitmap"
	"cg-go/src/colors"
	"cg-go/src/memory"
	"cg-go/src/vec"
	"image/color"
	"math"
)

type GeometricShape struct {
	Vertices        []vec.Vec2D
	ColorVertices   []color.RGBA
	TextureVertices []vec.Vec2D
}

func (s *GeometricShape) DrawBounds(mem memory.Memory) {
	if len(s.Vertices) < 3 {
		return
	}

	pi := s.Vertices[0]
	for i := 1; i < len(s.Vertices); i++ {
		pf := s.Vertices[i]
		bitmap.BresenhamLine(mem, pi, pf, colors.ColorWhite)
		pi = pf
	}

	pf := s.Vertices[0]
	bitmap.BresenhamLine(mem, pi, pf, colors.ColorWhite)

}

func (s *GeometricShape) WithColors(colors []color.RGBA) *GeometricShape {
	s.ColorVertices = colors
	return s
}

func (s *GeometricShape) WithTextureVertices(textureVertices []vec.Vec2D) *GeometricShape {
	s.TextureVertices = textureVertices
	return s
}

func NewRect(width, height float64, center vec.Vec2D) *GeometricShape {
	return &GeometricShape{
		Vertices: []vec.Vec2D{
			vec.NewVec2D(center.X-width, center.Y-height),
			vec.NewVec2D(center.X+width, center.Y-height),
			vec.NewVec2D(center.X+width, center.Y+height),
			vec.NewVec2D(center.X-width, center.Y+height),
		},
	}
}

func NewTriangle(base, height float64, center vec.Vec2D) *GeometricShape {
	return &GeometricShape{
		Vertices: []vec.Vec2D{
			vec.NewVec2D(center.X, center.Y-height/2),
			vec.NewVec2D(center.X+base/2, center.Y+height/2),
			vec.NewVec2D(center.X-base/2, center.Y+height/2),
		},
	}
}

func NewShape(vertices []vec.Vec2D) *GeometricShape {
	return &GeometricShape{Vertices: vertices}
}

func Copy(s *GeometricShape) *GeometricShape {
	v := make([]vec.Vec2D, len(s.Vertices))
	cv := make([]color.RGBA, len(s.ColorVertices))
	tv := make([]vec.Vec2D, len(s.TextureVertices))

	copy(v, s.Vertices)
	copy(cv, s.ColorVertices)
	copy(tv, s.TextureVertices)

	return &GeometricShape{
		Vertices:        v,
		ColorVertices:   cv,
		TextureVertices: tv,
	}
}

func (s *GeometricShape) Center() vec.Vec2D {
	var sumX, sumY float64

	for _, v := range s.Vertices {
		sumX += v.X
		sumY += v.Y
	}

	len := len(s.Vertices)
	centerX := sumX / float64(len)
	centerY := sumY / float64(len)

	return vec.NewVec2D(centerX, centerY)
}

func (s *GeometricShape) Apply(f func(s *GeometricShape)) {
	cp := Copy(s)
	f(cp)
}

func CreateCirclePolygon(center vec.Vec2D, radius float64, sides int) *GeometricShape {
	// Calcula o ângulo de cada lado do polígono
	angle := 2 * math.Pi / float64(sides)

	// Cria os vértices do polígono
	vertices := make([]vec.Vec2D, sides)
	for i := 0; i < sides; i++ {
		x := center.X + radius*math.Cos(float64(i)*angle)
		y := center.Y + radius*math.Sin(float64(i)*angle)
		vertices[i] = vec.NewVec2D(x, y)
	}

	// Cria e retorna o polígono
	return NewShape(vertices)
}
