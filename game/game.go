package game

import (
	"LudumDare/assets"

	"github.com/hajimehoshi/ebiten/v2"
)

type Game struct {
	bg *entity
	pl *entity
	W  int
	H  int
}

func New() *Game {
	i := ebiten.NewImageFromImage(assets.PL())
	w, h := i.Size()
	return &Game{
		bg: &entity{
			img: ebiten.NewImageFromImage(assets.BG()),
		},
		pl: &entity{
			img: ebiten.NewImageFromImage(assets.PL()),
			x:   w,
			y:   h,
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
	curledX := g.bg.x % g.W
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(float64(curledX), float64(g.bg.y))
	screen.DrawImage(g.bg.img, op)

	op = &ebiten.DrawImageOptions{}
	op.GeoM.Translate(float64(curledX+g.W), float64(g.bg.y))
	screen.DrawImage(g.bg.img, op)
}

func (g *Game) renderPlayer(screen *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(-float64(g.pl.x)/2, -float64(g.pl.y)/2)
	op.GeoM.Translate(
		float64(g.W)/2.0,
		float64(g.H)/2.0,
	)
	screen.DrawImage(g.pl.img, op)
}
