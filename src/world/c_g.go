package world

import (
	"cg-go/src/bitmap"
	"cg-go/src/colors"
	"cg-go/src/geo"
	"cg-go/src/memory"
	"cg-go/src/vec"
)

func DrawA(mem memory.Memory, pivot vec.Vec2D) {
	pol := &geo.GeometricShape{
		Vertices: []vec.Vec2D{pivot},
	}

	pol.Vertices = append(pol.Vertices, pivot.Plus(vec.NewVec2D(50, 0)))
	pol.Vertices = append(pol.Vertices, pivot.Plus(vec.NewVec2D(50, 50)))
	pol.Vertices = append(pol.Vertices, pivot.Plus(vec.NewVec2D(40, 50)))
	pol.Vertices = append(pol.Vertices, pivot.Plus(vec.NewVec2D(40, 30)))
	pol.Vertices = append(pol.Vertices, pivot.Plus(vec.NewVec2D(10, 30)))
	pol.Vertices = append(pol.Vertices, pivot.Plus(vec.NewVec2D(10, 50)))
	pol.Vertices = append(pol.Vertices, pivot.Plus(vec.NewVec2D(0, 50)))

	pol.DrawBounds(mem)
	geo.NewRect(15, 10, pivot.Plus(vec.NewVec2D(25, 15))).DrawBounds(mem)
	bitmap.FloodFill(mem, pivot.ScalarSum(1), colors.ColorIndigo, colors.ColorWhite)
}

func DrawC(mem memory.Memory, pivot vec.Vec2D) {
	pol := &geo.GeometricShape{
		Vertices: []vec.Vec2D{pivot},
	}

	pol.Vertices = append(pol.Vertices, pivot.Plus(vec.NewVec2D(50, 0)))
	pol.Vertices = append(pol.Vertices, pivot.Plus(vec.NewVec2D(50, 10)))
	pol.Vertices = append(pol.Vertices, pivot.Plus(vec.NewVec2D(20, 10)))
	pol.Vertices = append(pol.Vertices, pivot.Plus(vec.NewVec2D(20, 40)))
	pol.Vertices = append(pol.Vertices, pivot.Plus(vec.NewVec2D(50, 40)))
	pol.Vertices = append(pol.Vertices, pivot.Plus(vec.NewVec2D(50, 50)))
	pol.Vertices = append(pol.Vertices, pivot.Plus(vec.NewVec2D(0, 50)))

	pol.DrawBounds(mem)
	bitmap.FloodFill(mem, pivot.ScalarSum(1), colors.ColorIndigo, colors.ColorWhite)

}

func DrawG(mem memory.Memory, pivot vec.Vec2D) {
	pol := &geo.GeometricShape{
		Vertices: []vec.Vec2D{pivot},
	}

	pol.Vertices = append(pol.Vertices, pivot.Plus(vec.NewVec2D(50, 0)))
	pol.Vertices = append(pol.Vertices, pivot.Plus(vec.NewVec2D(50, 10)))
	pol.Vertices = append(pol.Vertices, pivot.Plus(vec.NewVec2D(20, 10)))
	pol.Vertices = append(pol.Vertices, pivot.Plus(vec.NewVec2D(20, 40)))
	pol.Vertices = append(pol.Vertices, pivot.Plus(vec.NewVec2D(40, 40)))
	pol.Vertices = append(pol.Vertices, pivot.Plus(vec.NewVec2D(40, 30)))
	pol.Vertices = append(pol.Vertices, pivot.Plus(vec.NewVec2D(35, 30)))
	pol.Vertices = append(pol.Vertices, pivot.Plus(vec.NewVec2D(35, 25)))
	pol.Vertices = append(pol.Vertices, pivot.Plus(vec.NewVec2D(50, 25)))
	pol.Vertices = append(pol.Vertices, pivot.Plus(vec.NewVec2D(50, 50)))
	pol.Vertices = append(pol.Vertices, pivot.Plus(vec.NewVec2D(0, 50)))

	pol.DrawBounds(mem)
	bitmap.FloodFill(mem, pivot.ScalarSum(1), colors.ColorIndigo, colors.ColorWhite)
}
