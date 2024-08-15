package world

import (
	"cg-go/src/colors"
	"cg-go/src/geo"
	"cg-go/src/memory"
	"cg-go/src/scan"
	"cg-go/src/tex"
	"cg-go/src/transform"
	"cg-go/src/vec"
	"cg-go/src/window"
	"image/color"
	"math/rand"
)

const Width, Height = 600, 400
const NumMeteors = 50
const MeteorsMinDist = 30
const WindowFactor = 2
const MeteorSize = 20

const WindowVelocity = 20

var meteorTex = tex.ReadImage("./resources/meteor.png")
var blackHoleTex = tex.ReadImage("./resources/Black_hole.png")
var terranTex = tex.ReadImage("./resources/Terran.png")
var lavaTex = tex.ReadImage("./resources/Lava.png")
var barenTex = tex.ReadImage("./resources/Baren.png")
var iceTex = tex.ReadImage("./resources/Ice.png")
var starTex = tex.ReadImage("./resources/star.png")
var gopherTex = tex.ReadImage("./resources/gopher-astronaut.png")

var Win = window.New(vec.NewVec2D(0, 0), vec.NewVec2D(Width*WindowFactor, Height*WindowFactor))
var Mem = memory.New(Width, Height)
var MainViewport = window.NewViewport(vec.Zeros(), vec.NewVec2D(Width, Height))
var MiniMap = window.NewViewport(vec.NewVec2D((Width-140), 20), vec.NewVec2D(Width-20, 100))

var stars = GenerateStars()
var meteors = GenerateMeteors()

var blackHole = geo.NewRect(100, 100, vec.NewVec2D(Width, Height)).WithTextureVertices([]vec.Vec2D{
	vec.Zeros(),
	vec.NewVec2D(1, 0),
	vec.Ones(),
	vec.NewVec2D(0, 1),
})

var terran = geo.NewRect(60, 60, vec.NewVec2D(150, 80)).WithTextureVertices([]vec.Vec2D{
	vec.Zeros(),
	vec.NewVec2D(1, 0),
	vec.Ones(),
	vec.NewVec2D(0, 1),
})

var lava = geo.NewRect(60, 60, vec.NewVec2D(250, 400)).WithTextureVertices([]vec.Vec2D{
	vec.Zeros(),
	vec.NewVec2D(1, 0),
	vec.Ones(),
	vec.NewVec2D(0, 1),
})

var baren = geo.NewRect(20, 20, lava.Center().ScalarSum(90)).WithTextureVertices([]vec.Vec2D{
	vec.Zeros(),
	vec.NewVec2D(1, 0),
	vec.Ones(),
	vec.NewVec2D(0, 1),
})

var ice = geo.NewRect(45, 45, vec.NewVec2D(900, 320)).WithTextureVertices([]vec.Vec2D{
	vec.Zeros(),
	vec.NewVec2D(1, 0),
	vec.Ones(),
	vec.NewVec2D(0, 1),
})

var shootingStar = geo.NewTriangle(20, 20, vec.NewVec2D(Width, Height)).WithColors([]color.RGBA{
	colors.ColorRed,
	colors.ColorYellow,
	colors.ColorYellow,
})

var gopher = geo.NewRect(20, 20, Win.Center()).
	WithTextureVertices([]vec.Vec2D{
		vec.NewVec2D(0, 0),
		vec.NewVec2D(1, 0),
		vec.NewVec2D(1, 1),
		vec.NewVec2D(0, 1),
	})

var pol = MakeSun(Win.Center().Sub(vec.NewVec2D(200, 200)), 100, 20)

var rect = geo.NewRect(200, 30, Win.Center().Plus(vec.NewVec2D(200, 200))).WithColors([]color.RGBA{
	colors.ColorBlue,
	colors.ColorRed,
	colors.ColorRed,
	colors.ColorBlue,
})

func MapObjectsToVP(vp *window.Viewport) {

	blackHole.Apply(func(s *geo.GeometricShape) {
		Win.MapPoints(s, vp)
		scan.ScanlineTexture(Mem, s, blackHoleTex)
	})

	terran.Apply(func(s *geo.GeometricShape) {
		Win.MapPoints(s, vp)
		scan.ScanlineTexture(Mem, s, terranTex)
	})

	lava.Apply(func(s *geo.GeometricShape) {
		Win.MapPoints(s, vp)
		scan.ScanlineTexture(Mem, s, lavaTex)
	})

	ice.Apply(func(s *geo.GeometricShape) {
		Win.MapPoints(s, vp)
		scan.ScanlineTexture(Mem, s, iceTex)
	})

	baren.Apply(func(s *geo.GeometricShape) {
		Win.MapPoints(s, vp)
		scan.ScanlineTexture(Mem, s, barenTex)
	})

	shootingStar.Apply(func(s *geo.GeometricShape) {
		Win.MapPoints(s, vp)
		scan.ScanlineGradient(Mem, s)
	})

	pol.Apply(func(s *geo.GeometricShape) {
		Win.MapPoints(s, vp)
		scan.ScanlineGradient(Mem, s)
	})

	rect.Apply(func(s *geo.GeometricShape) {
		Win.MapPoints(s, vp)
		scan.ScanlineGradient(Mem, s)
	})

	gopher.Apply(func(s *geo.GeometricShape) {
		Win.MapPoints(s, vp)
		scan.ScanlineTexture(Mem, s, gopherTex)
	})

}

func GenerateStars() (stars []*geo.GeometricShape) {
	for i := 0.0; i < Width*WindowFactor*2; i += 20 {
		star := geo.NewRect(5, 5, vec.NewVec2D(i, float64(rand.Intn(Height*WindowFactor*2)))).
			WithTextureVertices([]vec.Vec2D{
				vec.Zeros(),
				vec.NewVec2D(1, 0),
				vec.Ones(),
				vec.NewVec2D(0, 1),
			})
		stars = append(stars, star)
	}
	return
}

func DrawStars(mem memory.Memory) {
	for _, star := range stars {
		star.Apply(func(s *geo.GeometricShape) {
			Win.MapPoints(s, MainViewport)
			scan.ScanlineTexture(Mem, s, starTex)
		})
	}
}

func DrawMeteors(mem memory.Memory) {
	for _, meteor := range meteors {
		meteor.Apply(func(s *geo.GeometricShape) {
			Win.MapPoints(s, MainViewport)
			scan.ScanlineTexture(Mem, s, meteorTex)
		})

		meteor.Apply(func(s *geo.GeometricShape) {
			Win.MapPoints(s, MiniMap)
			scan.ScanlineTexture(Mem, s, meteorTex)
		})
	}
}

func RotateMeteors() {
	for i := range meteors {
		if i%2 == 0 {
			transform.RotateVerticesOnPivot(-1, meteors[i].Center(), meteors[i])
		} else {
			transform.RotateVerticesOnPivot(1, meteors[i].Center(), meteors[i])
		}

	}
}

func MeteorsPositions() []vec.Vec2D {

	points := make([]vec.Vec2D, NumMeteors)

	for i := 0; i < NumMeteors; i++ {
		var x, y float64
		isValid := false

		for !isValid {
			isValid = true
			x = rand.Float64() * Width * WindowFactor * 2
			y = rand.Float64() * Height * WindowFactor * 2

			for j := 0; j < i; j++ {
				if vec.Distance(points[j], vec.NewVec2D(x, y)) < MeteorsMinDist {
					isValid = false
					break
				}
			}
		}

		points[i] = vec.NewVec2D(x, y)
	}

	return points
}

func GenerateMeteors() (meteors []*geo.GeometricShape) {
	for _, pos := range MeteorsPositions() {
		meteors = append(meteors, geo.NewRect(MeteorSize, MeteorSize, pos).WithTextureVertices([]vec.Vec2D{
			vec.Zeros(),
			vec.NewVec2D(1, 0),
			vec.Ones(),
			vec.NewVec2D(0, 1),
		}))
	}
	return
}

func MakeSun(center vec.Vec2D, radius float64, sides int) *geo.GeometricShape {

	c := make([]color.RGBA, sides)

	for i := 0; i < sides; i++ {

		if i%2 == 0 {
			c[i] = colors.ColorYellow
		} else {
			c[i] = colors.ColorRed
		}

	}

	return geo.CreateCirclePolygon(Win.Center().Sub(vec.NewVec2D(200, 200)), 100, 12).
		WithColors(c)
}
