package game

import (
	"LudumDare/assets"

	"github.com/hajimehoshi/ebiten/v2"
)

type Game struct {
	bg *entity
	W  int
	H  int
}

func New() *Game {
	return &Game{
		bg: &entity{
			img: ebiten.NewImageFromImage(assets.Background()),
		},
	}
}

func (g *Game) Update() error {
	g.updateBackground()
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	g.renderBackground(screen)
	g.renderPlayer(screen)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return g.W, g.H
}

func (g *Game) updateBackground() {
	g.bg.x -= 2
}

func (g *Game) renderBackground(screen *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(float64(g.bg.x), float64(g.bg.y))
	screen.DrawImage(g.bg.img, op)
}

func (g *Game) renderPlayer(screen *ebiten.Image) {

}
