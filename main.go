package main

import (
	"cg-go/src/screen"
	"cg-go/src/world"
)

func main() {

	screen.New().
		SetWidth(world.Mem.Width()).
		SetHeight(world.Mem.Height()).
		SetTitle("CG").
		SetOnUpdate(world.Update).
		Build().
		Run()

}
