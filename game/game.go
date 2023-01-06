package game

import (
	"github.com/hajimehoshi/ebiten/v2"
)

type Game struct {
	W int
	H int
}

func New() *Game {
	return &Game{}
}

func (g *Game) Update() error {
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	g.renderBackground(screen)
	g.renderPlayer(screen)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return g.W, g.H
}

func (g *Game) renderBackground(screen *ebiten.Image) {
}

func (g *Game) renderPlayer(screen *ebiten.Image) {

}
