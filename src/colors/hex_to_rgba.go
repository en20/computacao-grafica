package colors

import "image/color"

func HexToRGBA(hex uint32) color.RGBA {
	a := uint8(hex >> 24 & 0xFF)
	r := uint8(hex >> 16 & 0xFF)
	g := uint8(hex >> 8 & 0xFF)
	b := uint8(hex & 0xFF)

	return color.RGBA{r, g, b, a}
}
