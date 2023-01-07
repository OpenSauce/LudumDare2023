package entities

import (
	"github.com/hajimehoshi/ebiten/v2"
	"golang.org/x/image/math/f64"
)

type Entity struct {
	Pos       f64.Vec2
	W         int
	H         int
	Img       *ebiten.Image
	Rotation  int
	Direction f64.Vec2
	Scale     float64

	HP int
}
