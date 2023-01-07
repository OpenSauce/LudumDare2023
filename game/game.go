package game

import (
	"LudumDare/assets"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/audio"
	"github.com/hajimehoshi/ebiten/v2/audio/mp3"
)

type Game struct {
	bg *entity
	pl *entity
	W  int
	H  int
}

func New(sW, sH int) *Game {
	audioContext := audio.NewContext(44100)

	d, err := mp3.DecodeWithSampleRate(44100, assets.LoadMusic())
	if err != nil {
		log.Fatal("error loading music")
	}

	p, err := audioContext.NewPlayer(d)
	if err != nil {
		log.Fatal("error loading music")
	}

	p.SetVolume(0.2)

	p.Play()

	i := ebiten.NewImageFromImage(assets.Turtle())
	w, h := i.Size()
	return &Game{
		W: sW,
		H: sH,
		bg: &entity{
			img: ebiten.NewImageFromImage(assets.Background()),
		},
		pl: &entity{
			img: ebiten.NewImageFromImage(assets.Turtle()),
			w:   w,
			h:   h,
			x:   sW / 4,
			y:   sH / 4,
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
		g.pl.x += 1
	}

	if ebiten.IsKeyPressed(ebiten.KeyArrowLeft) {
		g.pl.x -= 1
	}

	if ebiten.IsKeyPressed(ebiten.KeyArrowUp) {
		g.pl.y -= 1
	}

	if ebiten.IsKeyPressed(ebiten.KeyArrowDown) {
		g.pl.y += 1
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
	op.GeoM.Scale(2, 2)
	screen.DrawImage(g.pl.img, op)
}
