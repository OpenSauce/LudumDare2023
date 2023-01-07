package game

import (
	"LudumDare/assets"
	"LudumDare/entities"
	"LudumDare/util"
	"fmt"
	"image/color"
	"log"
	"math"
	"math/rand"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/examples/resources/fonts"
	"github.com/hajimehoshi/ebiten/v2/text"
	"golang.org/x/image/font"
	"golang.org/x/image/font/opentype"
	"golang.org/x/image/math/f64"
)

var (
	debug = true

	actionTime         = 15
	turretsActionTimes = map[*entities.Entity]int{}
	projectileLifeTime = 200

	gameFont font.Face
)

type MainScreen struct {
	count       int
	w           int
	h           int
	bg          *entities.Entity
	pl          *entities.Entity
	turrets     map[*entities.Entity]f64.Vec2
	projectiles map[*entities.Entity]int
	enemies     map[*entities.Entity]struct{}
}

func NewMainScreen(sW, sH int) *MainScreen {
	tt, err := opentype.Parse(fonts.MPlus1pRegular_ttf)
	if err != nil {
		log.Fatal(err)
	}

	const dpi = 72
	gameFont, err = opentype.NewFace(tt, &opentype.FaceOptions{
		Size:    24,
		DPI:     dpi,
		Hinting: font.HintingFull,
	})
	if err != nil {
		log.Fatal(err)
	}

	i := ebiten.NewImageFromImage(assets.Turtle())
	w, h := i.Size()
	ms := &MainScreen{
		w: sW,
		h: sH,
		bg: &entities.Entity{
			Img: ebiten.NewImageFromImage(assets.Background()),
		},
		pl: &entities.Entity{
			Img:   ebiten.NewImageFromImage(i),
			HP:    100,
			W:     w,
			H:     h,
			Pos:   f64.Vec2{float64(sW / 4), float64(sH / 4)},
			Scale: 2,
		},
		turrets: map[*entities.Entity]f64.Vec2{
			{
				Img:   ebiten.NewImageFromImage(assets.Turret()),
				Scale: 1,
			}: {10, +22},
			{
				Img:   ebiten.NewImageFromImage(assets.Turret()),
				Scale: 1,
			}: {10, -28},
		},
		projectiles: map[*entities.Entity]int{},
		enemies:     map[*entities.Entity]struct{}{},
	}

	for t := range ms.turrets {
		turretsActionTimes[t] = 0
	}

	return ms
}

func (m *MainScreen) Update() error {
	m.count = (m.count + 1) % 60

	if ebiten.IsKeyPressed(ebiten.KeyEscape) {
		return fmt.Errorf("force closed game")
	}

	m.updatePlayer()
	m.updateBackground()
	m.updateTurrets()
	m.updateProjectiles()
	m.updateEnemies()
	return nil
}

func (m *MainScreen) Draw(screen *ebiten.Image) {
	m.renderBackground(screen)
	m.renderPlayer(screen)
	m.renderTurrets(screen)
	m.renderProjectiles(screen)
	m.renderEnemies(screen)

	if debug {
		m.debug(screen)
	}
}

func (m *MainScreen) updatePlayer() {
	// TODO: Can't travel out of bounds
	if ebiten.IsKeyPressed(ebiten.KeyD) || ebiten.IsKeyPressed(ebiten.KeyE) {
		m.pl.Pos[0] += 2
	}

	if ebiten.IsKeyPressed(ebiten.KeyA) {
		m.pl.Pos[0] -= 2
	}

	if ebiten.IsKeyPressed(ebiten.KeyW) || ebiten.IsKeyPressed(ebiten.KeyComma) {
		m.pl.Pos[1] -= 2
	}

	if ebiten.IsKeyPressed(ebiten.KeyS) || ebiten.IsKeyPressed(ebiten.KeyO) {
		m.pl.Pos[1] += 2
	}
}

func (m *MainScreen) updateBackground() {
	m.bg.Pos[0] -= 10
	m.bg.Pos[0] = float64(int(m.bg.Pos[0]) % m.w)
}

func (m *MainScreen) updateTurrets() {
	x, y := ebiten.CursorPosition()

	for t, offset := range m.turrets {
		t.Pos[0] = m.pl.Pos[0] + offset[0]
		t.Pos[1] = m.pl.Pos[1] + offset[1]

		xf := float64(x) - t.Pos[0]*2
		yf := float64(y) - t.Pos[1]*2

		angle := math.Atan2(yf, xf) * (180 / math.Pi)

		t.Rotation = int(angle) + 90

		turretsActionTimes[t]++
		if turretsActionTimes[t] >= actionTime && ebiten.IsMouseButtonPressed(ebiten.MouseButtonLeft) {
			turretsActionTimes[t] = 0
			xdir := (t.Pos[0] * 2) - float64(x)
			ydir := (t.Pos[1] * 2) - float64(y)

			m.projectiles[&entities.Entity{
				Pos:       f64.Vec2{t.Pos[0] * 2, t.Pos[1] * 2},
				Img:       ebiten.NewImageFromImage(assets.Projectile()),
				Rotation:  t.Rotation,
				Direction: f64.Vec2{xdir, ydir},
				Scale:     1,
			}] = projectileLifeTime
		}
	}
}

func (m *MainScreen) updateEnemies() {
	if m.count == 0 {
		i := ebiten.NewImageFromImage(assets.Enemy())
		w, h := i.Size()
		startingY := rand.Intn(m.h - h)
		m.enemies[&entities.Entity{
			Img:   ebiten.NewImageFromImage(i),
			Scale: 1,
			W:     w,
			H:     h,
			Pos:   f64.Vec2{float64(m.w + w), float64(startingY)},
		}] = struct{}{}
	}

	for enemy := range m.enemies {
		if util.DoesCollide(*enemy, *m.pl) {
			m.pl.HP += 20
			delete(m.enemies, enemy)
			continue
		}
		enemy.Pos[0] -= 5
	}
}

func (m *MainScreen) updateProjectiles() {
	deadProj := []*entities.Entity{}

	for p, lt := range m.projectiles {
		if lt <= 0 {
			deadProj = append(deadProj, p)
			continue
		}
		m.projectiles[p] -= 1

		mag := getMag(p.Direction)

		p.Pos[0] -= p.Direction[0] / mag * 10
		p.Pos[1] -= p.Direction[1] / mag * 10
	}

	for _, dp := range deadProj {
		delete(m.projectiles, dp)
		delete(turretsActionTimes, dp)
	}
}

func (m *MainScreen) renderBackground(screen *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(m.bg.Pos[0], m.bg.Pos[1])
	screen.DrawImage(m.bg.Img, op)

	op = &ebiten.DrawImageOptions{}
	op.GeoM.Translate(m.bg.Pos[0]+float64(m.w), m.bg.Pos[1])
	screen.DrawImage(m.bg.Img, op)
}

func (m *MainScreen) renderPlayer(screen *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(m.pl.Pos[0]-float64(m.pl.W)/2, m.pl.Pos[1]-float64(m.pl.H)/2)
	op.GeoM.Scale(m.pl.Scale, m.pl.Scale)
	screen.DrawImage(m.pl.Img, op)
}

func (m *MainScreen) renderTurrets(screen *ebiten.Image) {
	for t := range m.turrets {
		w, h := t.Img.Size()

		op := &ebiten.DrawImageOptions{}

		op.GeoM.Translate(-float64(w)/2, -float64(h)/2)
		op.GeoM.Rotate(float64(t.Rotation%360) * 2 * math.Pi / 360)

		op.GeoM.Translate(t.Pos[0]*2, t.Pos[1]*2)
		op.GeoM.Scale(t.Scale, t.Scale)
		screen.DrawImage(t.Img, op)
	}

}

func (m *MainScreen) renderEnemies(screen *ebiten.Image) {
	for e := range m.enemies {
		w, h := e.Img.Size()

		op := &ebiten.DrawImageOptions{}
		op.GeoM.Translate(-float64(w)/2, -float64(h)/2)
		op.GeoM.Translate(e.Pos[0], e.Pos[1])
		op.GeoM.Scale(e.Scale, e.Scale)
		screen.DrawImage(e.Img, op)
	}

}

func (m *MainScreen) renderProjectiles(screen *ebiten.Image) {
	for p := range m.projectiles {
		w, h := p.Img.Size()
		op := &ebiten.DrawImageOptions{}

		op.GeoM.Translate(-float64(w)/2, -float64(h)/2)

		op.GeoM.Rotate(float64(p.Rotation%360) * 2 * math.Pi / 360)

		op.GeoM.Translate(p.Pos[0], p.Pos[1])
		op.GeoM.Scale(p.Scale, p.Scale)

		screen.DrawImage(p.Img, op)
	}
}

func (m *MainScreen) debug(screen *ebiten.Image) {
	text.Draw(screen, fmt.Sprintf("Pos: [%v, %v] HP: %v",
		m.pl.Pos[0], m.pl.Pos[1], m.pl.HP), gameFont, 20, 20, color.White)
}

func getMag(vec2 f64.Vec2) float64 {
	return math.Sqrt(vec2[0]*vec2[0] + vec2[1]*vec2[1])
}
