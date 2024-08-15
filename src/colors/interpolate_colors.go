package colors

import (
	"image/color"
	"math"
)

func InterpolateColors(startColor color.RGBA, endColor color.RGBA, t float64) color.RGBA {

	r := uint8(math.Round(float64(startColor.R) + float64(float64(endColor.R)-float64(startColor.R))*t))
	g := uint8(math.Round(float64(startColor.G) + float64(float64(endColor.G)-float64(startColor.G))*t))
	b := uint8(math.Round(float64(startColor.B) + float64(float64(endColor.B)-float64(startColor.B))*t))
	a := uint8(math.Round(float64(startColor.A) + float64(float64(endColor.A)-float64(startColor.A))*t))
	return color.RGBA{r, g, b, a}
}
