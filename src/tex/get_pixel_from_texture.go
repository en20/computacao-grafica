package tex

import (
	"image/color"
	"math"
)

func GetPixelFromTexture(image [][]color.RGBA, x, y float64) color.RGBA {
	x, y = math.Max(math.Min(x, 1), 0), math.Max(math.Min(y, 1), 0)

	rows, columns := len(image[0]), len(image)

	imgX := math.Round(float64(rows-1) * x)
	imgY := math.Round(float64(columns-1) * y)

	return image[int(math.Round(imgY))][int(math.Round(imgX))]
}
