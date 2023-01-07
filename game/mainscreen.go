package game

import (
	"LudumDare/assets"
	"math"

	"github.com/hajimehoshi/ebiten/v2"
	"golang.org/x/image/math/f64"
)

var (
	actionTime         = 15
	turretsActionTimes = map[*entity]int{}
	projectileLifeTime = 200
)

type MainScreen struct {
	w           int
	h           int
	bg          *entity
	pl          *entity
	turrets     map[*entity]f64.Vec2
	projectiles map[*entity]int
}

func NewMainScreen(sW, sH int) *MainScreen {
	i := ebiten.NewImageFromImage(assets.Turtle())
	w, h := i.Size()
	ms := &MainScreen{
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
		projectiles: map[*entity]int{},
	}

	for t := range ms.turrets {
		turretsActionTimes[t] = 0
	}

	return ms
}

func (m *MainScreen) Update() error {
	m.updatePlayer()
	m.updateBackground()
	m.updateTurrets()
	m.updateProjectiles()
	return nil
}

func (m *MainScreen) Draw(screen *ebiten.Image) {
	m.renderBackground(screen)
	m.renderPlayer(screen)
	m.renderTurrets(screen)
	m.renderProjectiles(screen)
}

func (m *MainScreen) updatePlayer() {
	if ebiten.IsKeyPressed(ebiten.KeyD) {
		m.pl.pos[0] += 2
	}

	if ebiten.IsKeyPressed(ebiten.KeyA) {
		m.pl.pos[0] -= 2
	}

	if ebiten.IsKeyPressed(ebiten.KeyW) {
		m.pl.pos[1] -= 2
	}

	if ebiten.IsKeyPressed(ebiten.KeyS) {
		m.pl.pos[1] += 2
	}
}

func (m *MainScreen) updateBackground() {
	m.bg.pos[0] -= 10
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

		turretsActionTimes[t]++
		if turretsActionTimes[t] >= actionTime && ebiten.IsMouseButtonPressed(ebiten.MouseButtonLeft) {
			turretsActionTimes[t] = 0
			xdir := (t.pos[0] * 2) - float64(x)
			ydir := (t.pos[1] * 2) - float64(y)

			m.projectiles[&entity{
				pos:       f64.Vec2{t.pos[0] * 2, t.pos[1] * 2},
				img:       ebiten.NewImageFromImage(assets.Projectile()),
				rotation:  t.rotation,
				direction: f64.Vec2{xdir, ydir},
			}] = projectileLifeTime
		}
	}
}

func (m *MainScreen) updateProjectiles() {
	deadProj := []*entity{}

	for p, lt := range m.projectiles {
		if lt <= 0 {
			deadProj = append(deadProj, p)
			continue
		}
		m.projectiles[p] -= 1

		mag := getMag(p.direction)

		p.pos[0] -= p.direction[0] / mag * 10
		p.pos[1] -= p.direction[1] / mag * 10
	}

	for _, dp := range deadProj {
		delete(m.projectiles, dp)
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
	for t := range m.turrets {
		w, h := t.img.Size()

		op := &ebiten.DrawImageOptions{}

		op.GeoM.Translate(-float64(w)/2, -float64(h)/2)
		op.GeoM.Rotate(float64(t.rotation%360) * 2 * math.Pi / 360)

		op.GeoM.Translate(t.pos[0]*2, t.pos[1]*2)
		op.GeoM.Scale(1, 1)
		screen.DrawImage(t.img, op)
	}

}

func (m *MainScreen) renderProjectiles(screen *ebiten.Image) {
	for p := range m.projectiles {
		w, h := p.img.Size()
		op := &ebiten.DrawImageOptions{}

		op.GeoM.Translate(-float64(w)/2, -float64(h)/2)

		op.GeoM.Rotate(float64(p.rotation%360) * 2 * math.Pi / 360)

		op.GeoM.Translate(p.pos[0], p.pos[1])
		op.GeoM.Scale(1, 1)

		screen.DrawImage(p.img, op)
	}
}

func getMag(vec2 f64.Vec2) float64 {
	return math.Sqrt(vec2[0]*vec2[0] + vec2[1]*vec2[1])
}
