package main

import (
	"github.com/hajimehoshi/ebiten/v2"

	game "github.com/byrdbrandon6/astro/scenes"
	"github.com/byrdbrandon6/astro/vars"
)

func main() {
	g := game.NewGame()
	ebiten.SetWindowSize(vars.ScreenWidth*vars.Scale, vars.ScreenHeight*vars.Scale)
	ebiten.SetWindowTitle("Astro")

	err := ebiten.RunGame(g)
	if err != nil {
		panic(err)
	}
}
