package colors

const (
	Black  = 0xff000000
	White  = 0xffffffff
	Silver = 0xffc6c6c6
	Gray   = 0xff808080
	Red    = 0xffff3b30
	Orange = 0xffff9500
	Yellow = 0xffffcc00
	Green  = 0xff4cd964
	Teal   = 0xff5ac8fa
	Blue   = 0xff007aff
	Indigo = 0xff5856d6
	Purple = 0xff1d224c

	Pink = 0xffff2d55
)

var (
	ColorBlack  = HexToRGBA(Black)
	ColorWhite  = HexToRGBA(White)
	ColorSilver = HexToRGBA(Silver)
	ColorRed    = HexToRGBA(Red)
	ColorOrange = HexToRGBA(Orange)
	ColorYellow = HexToRGBA(Yellow)
	ColorGreen  = HexToRGBA(Green)
	ColorTeal   = HexToRGBA(Teal)
	ColorBlue   = HexToRGBA(Blue)
	ColorIndigo = HexToRGBA(Indigo)
	ColorPurple = HexToRGBA(Purple)
	ColorPink   = HexToRGBA(Pink)
)
