package game

import (
	"LudumDare/assets"
	"math"

	"github.com/hajimehoshi/ebiten/v2"
	"golang.org/x/image/math/f64"
)

type MainScreen struct {
	w       int
	h       int
	bg      *entity
	pl      *entity
	turrets map[*entity]f64.Vec2
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
			pos: f64.Vec2{float64(sW / 4), float64(sH / 4)},
		},
		turrets: map[*entity]f64.Vec2{
			{
				img: ebiten.NewImageFromImage(assets.Turret()),
			}: {10, +22},
			{
				img: ebiten.NewImageFromImage(assets.Turret()),
			}: {10, -28},
		},
	}
}

func (m *MainScreen) Update() error {
	m.updatePlayer()
	m.updateBackground()
	m.updateTurrets()
	return nil
}

func (m *MainScreen) Draw(screen *ebiten.Image) {
	m.renderBackground(screen)
	m.renderPlayer(screen)
	m.renderTurrets(screen)
}

func (m *MainScreen) updatePlayer() {
	if ebiten.IsKeyPressed(ebiten.KeyD) {
		m.pl.pos[0] += 1
	}

	if ebiten.IsKeyPressed(ebiten.KeyA) {
		m.pl.pos[0] -= 1
	}

	if ebiten.IsKeyPressed(ebiten.KeyW) {
		m.pl.pos[1] -= 1
	}

	if ebiten.IsKeyPressed(ebiten.KeyS) {
		m.pl.pos[1] += 1
	}
}

func (m *MainScreen) updateBackground() {
	m.bg.pos[0] -= 2
	m.bg.pos[0] = float64(int(m.bg.pos[0]) % m.w)
}

func (m *MainScreen) updateTurrets() {
	x, y := ebiten.CursorPosition()

	for t, offset := range m.turrets {

		t.pos[0] = m.pl.pos[0] + offset[0]
		t.pos[1] = m.pl.pos[1] + offset[1]

		xf := float64(x) - t.pos[0]*2
		yf := float64(y) - t.pos[1]*2

		angle := math.Atan2(yf, xf) * (180 / math.Pi)

		t.rotation = int(angle) + 90

		// currentActionTime++
		// if currentActionTime >= actionTime && ebiten.IsMouseButtonPressed(ebiten.MouseButtonLeft) {
		// 	currentActionTime = 0
		// 	xdir := t.pos[0] - float64(x)
		// 	ydir := t.pos[1] - float64(y)

		// 	g.projectiles.New(ScreenWidth/2, ScreenHeight/2, f64.Vec2{xdir, ydir}, t.rotation)
		// }

	}
}

func (m *MainScreen) renderBackground(screen *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(m.bg.pos[0], m.bg.pos[1])
	screen.DrawImage(m.bg.img, op)

	op = &ebiten.DrawImageOptions{}
	op.GeoM.Translate(m.bg.pos[0]+float64(m.w), m.bg.pos[1])
	screen.DrawImage(m.bg.img, op)
}

func (m *MainScreen) renderPlayer(screen *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(m.pl.pos[0]-float64(m.pl.w)/2, m.pl.pos[1]-float64(m.pl.h)/2)
	op.GeoM.Scale(2, 2)
	screen.DrawImage(m.pl.img, op)
}

func (m *MainScreen) renderTurrets(screen *ebiten.Image) {
	for t, _ := range m.turrets {
		w, h := t.img.Size()

		op := &ebiten.DrawImageOptions{}

		op.GeoM.Translate(-float64(w)/2, -float64(h)/2)
		op.GeoM.Rotate(float64(t.rotation%360) * 2 * math.Pi / 360)

		op.GeoM.Translate(t.pos[0]*2, t.pos[1]*2)
		op.GeoM.Scale(1, 1)
		screen.DrawImage(t.img, op)
	}

}
