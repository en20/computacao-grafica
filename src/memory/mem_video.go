package memory

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
)

type memVideo struct {
	width  int
	height int
	mem    []color.RGBA
}

func New(width, height int) *memVideo {
	return &memVideo{
		width:  width,
		height: height,
		mem:    make([]color.RGBA, width*height),
	}
}

func (m *memVideo) SetPixel(x, y int, color color.RGBA) {
	// Para aplicar a transparência.
	if color.A == 0 {
		return
	}

	if x >= m.width || x < 0 || y >= m.height || y < 0 {
		return
	}

	idx := m.calcPosition(x, y)

	if idx >= len(m.mem) {
		return
	}

	m.mem[idx] = color
}

func (m *memVideo) GetPixel(x, y int) color.RGBA {
	if x >= m.width || x < 0 || y >= m.height || y < 0 {
		return color.RGBA{}
	}

	idx := m.calcPosition(x, y)

	if idx >= len(m.mem) {
		return color.RGBA{}
	}

	return m.mem[idx]
}

func (m *memVideo) Draw(ctx *ebiten.Image) {
	for i := 0; i < m.width; i++ {
		for j := 0; j < m.height; j++ {
			ctx.Set(i, j, m.GetPixel(i, j))
		}
	}
}

func (m *memVideo) calcPosition(x, y int) int {
	return y*m.width + x
}

func (m *memVideo) Clear(color color.RGBA) {
	for i := range m.mem {
		m.mem[i] = color
	}

}

func (m *memVideo) Width() int {
	return m.width
}

func (m *memVideo) Height() int {
	return m.height
}
