package bitmap

import (
	"cg-go/src/memory"
	"cg-go/src/vec"
	"image/color"
)

type point struct {
	X, Y int
}

func FloodFill(mem memory.Memory, startPoint vec.Vec2D, colorBg, colorBorder color.RGBA) {
	stack := []point{{int(startPoint.X), int(startPoint.Y)}}
	visited := make(map[point]bool)

	for len(stack) > 0 {
		// Desempilha um ponto da pilha
		currPoint := stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		// Verifica se o ponto já foi visitado
		if visited[currPoint] {
			continue
		}

		// Obtém a cor do pixel atual
		currColor := mem.GetPixel(currPoint.X, currPoint.Y)

		// Verifica se o pixel já foi preenchido ou se é uma borda
		if currColor == colorBg || currColor == colorBorder {
			continue
		}

		// Preenche o pixel com a nova cor
		mem.SetPixel(currPoint.X, currPoint.Y, colorBg)

		// Marca o ponto como visitado
		visited[currPoint] = true

		// Empilha os pontos vizinhos na pilha
		stack = append(stack, getNeighbors(currPoint)...)
	}
}

func getNeighbors(p point) []point {
	return []point{
		{p.X + 1, p.Y},
		{p.X - 1, p.Y},
		{p.X, p.Y + 1},
		{p.X, p.Y - 1},
	}
}
