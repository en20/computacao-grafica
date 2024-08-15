package screen

import (
	"log"

	"github.com/hajimehoshi/ebiten/v2"
)

type Screen struct {
	update        func(screen *ebiten.Image)
	title         string
	width, height int
}

func (s *Screen) Update() error {
	return nil
}

func (s *Screen) Draw(screen *ebiten.Image) {
	s.update(screen)
}

func (s *Screen) Layout(outsideWidth, outsideHeight int) (int, int) {
	return s.width, s.height
}

func (s *Screen) SetWidth(width int) *Screen {
	s.width = width
	return s
}

func (s *Screen) SetHeight(height int) *Screen {
	s.height = height
	return s
}

func (s *Screen) SetOnUpdate(update func(screen *ebiten.Image)) *Screen {
	s.update = update
	return s
}

func (s *Screen) SetTitle(title string) *Screen {
	s.title = title
	return s
}

func (s *Screen) Build() *Screen {
	ebiten.SetWindowTitle(s.title)
	ebiten.SetWindowSize(s.width, s.height)
	return s
}

func New() *Screen {
	return &Screen{}
}

func (s *Screen) Run() {
	if err := ebiten.RunGameWithOptions(s, &ebiten.RunGameOptions{}); err != nil {
		log.Fatal(err)
	}
}
