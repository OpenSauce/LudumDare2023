package main

import (
	"LudumDare/game"

	"github.com/hajimehoshi/ebiten/v2"
)

const gameName = "LudumDare"

func main() {
	ebiten.SetFullscreen(true)
	ebiten.SetWindowTitle(gameName)
	w, h := ebiten.ScreenSizeInFullscreen()
	game := game.New(w, h)
	if err := ebiten.RunGame(game); err != nil {
		panic(err)
	}
}
