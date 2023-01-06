package main

import (
	"LudumDare/game"

	"github.com/hajimehoshi/ebiten/v2"
)

const gameName = "LudumDare"

func main() {
	ebiten.SetFullscreen(true)
	ebiten.SetWindowTitle(gameName)
	game := game.New()
	game.W, game.H = ebiten.ScreenSizeInFullscreen()
	if err := ebiten.RunGame(game); err != nil {
		panic(err)
	}
}
