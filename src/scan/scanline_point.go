package scan

import "image/color"

type ScanlinePoint struct {
	X int
	T float64
}

type ScanlinePointGradient struct {
	ScanlinePoint
	Color color.RGBA
}

type ScanlinePointTexture struct {
	X, Y   int
	Tx, Ty float64
}

func NewScanlinePoint(xi int, t float64) ScanlinePoint {
	return ScanlinePoint{xi, t}
}

func NewScanlinePointGradient(xi int, t float64, color color.RGBA) ScanlinePointGradient {
	return ScanlinePointGradient{NewScanlinePoint(xi, t), color}
}

func NewScanlinePointTexture(x, y int, tx, ty float64) ScanlinePointTexture {
	return ScanlinePointTexture{x, y, tx, ty}
}
