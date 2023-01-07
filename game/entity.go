package game

import (
	"github.com/hajimehoshi/ebiten/v2"
	"golang.org/x/image/math/f64"
)

type entity struct {
	pos      f64.Vec2
	w        int
	h        int
	img      *ebiten.Image
	rotation int
}
