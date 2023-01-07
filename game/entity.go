package game

import "github.com/hajimehoshi/ebiten/v2"

type entity struct {
	x   int
	y   int
	w   int
	h   int
	img *ebiten.Image
}
