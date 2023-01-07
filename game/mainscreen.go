package game

import (
	"LudumDare/assets"

	"github.com/hajimehoshi/ebiten/v2"
)

type MainScreen struct {
	w  int
	h  int
	bg *entity
	pl *entity
}

func NewMainScreen(sW, sH int) *MainScreen {
	i := ebiten.NewImageFromImage(assets.Turtle())
	w, h := i.Size()
	return &MainScreen{
		w: sW,
		h: sH,
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

func (m *MainScreen) Update() error {
	m.updatePlayer()
	m.updateBackground()
	return nil
}

func (m *MainScreen) Draw(screen *ebiten.Image) {
	m.renderBackground(screen)
	m.renderPlayer(screen)
}

func (m *MainScreen) updatePlayer() {
	if ebiten.IsKeyPressed(ebiten.KeyArrowRight) {
		m.pl.x += 1
	}

	if ebiten.IsKeyPressed(ebiten.KeyArrowLeft) {
		m.pl.x -= 1
	}

	if ebiten.IsKeyPressed(ebiten.KeyArrowUp) {
		m.pl.y -= 1
	}

	if ebiten.IsKeyPressed(ebiten.KeyArrowDown) {
		m.pl.y += 1
	}
}

func (m *MainScreen) updateBackground() {
	m.bg.x -= 2
	m.bg.x %= m.w
}

func (m *MainScreen) renderBackground(screen *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(float64(m.bg.x), float64(m.bg.y))
	screen.DrawImage(m.bg.img, op)

	op = &ebiten.DrawImageOptions{}
	op.GeoM.Translate(float64(m.bg.x+m.w), float64(m.bg.y))
	screen.DrawImage(m.bg.img, op)
}

func (m *MainScreen) renderPlayer(screen *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(float64(m.pl.x)-float64(m.pl.w)/2, float64(m.pl.y)-float64(m.pl.h)/2)
	op.GeoM.Scale(2, 2)
	screen.DrawImage(m.pl.img, op)
}
