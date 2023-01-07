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

func New(sW, sH int) *Game {
	i := ebiten.NewImageFromImage(assets.PL())
	w, h := i.Size()
	return &Game{
		W: sW,
		H: sH,
		bg: &entity{
			img: ebiten.NewImageFromImage(assets.BG()),
		},
		pl: &entity{
			img: ebiten.NewImageFromImage(assets.PL()),
			w:   w,
			h:   h,
			x:   sW / 4,
			y:   sH / 2,
		},
	}
}

func (g *Game) Update() error {
	g.updateBackground()
	g.updatePlayer()
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	g.renderBackground(screen)
	g.renderPlayer(screen)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return g.W, g.H
}

func (g *Game) updatePlayer() {
	if ebiten.IsKeyPressed(ebiten.KeyArrowRight) {
		g.pl.x += 3
	}

	if ebiten.IsKeyPressed(ebiten.KeyArrowLeft) {
		g.pl.x -= 3
	}

	if ebiten.IsKeyPressed(ebiten.KeyArrowUp) {
		g.pl.y -= 3
	}

	if ebiten.IsKeyPressed(ebiten.KeyArrowDown) {
		g.pl.y += 3
	}
}

func (g *Game) updateBackground() {
	g.bg.x -= 2
	g.bg.x %= g.W
}

func (g *Game) renderBackground(screen *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(float64(g.bg.x), float64(g.bg.y))
	screen.DrawImage(g.bg.img, op)

	op = &ebiten.DrawImageOptions{}
	op.GeoM.Translate(float64(g.bg.x+g.W), float64(g.bg.y))
	screen.DrawImage(g.bg.img, op)
}

func (g *Game) renderPlayer(screen *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(float64(g.pl.x)-float64(g.pl.w)/2, float64(g.pl.y)-float64(g.pl.h)/2)
	screen.DrawImage(g.pl.img, op)
}
