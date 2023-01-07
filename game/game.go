package game

import (
	"LudumDare/assets"
	"log"
	"math"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/audio"
	"github.com/hajimehoshi/ebiten/v2/audio/mp3"
	"golang.org/x/image/math/f64"
)

type Game struct {
	bg      *entity
	pl      *entity
	turrets map[*entity]f64.Vec2
	W       int
	H       int
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

func (g *Game) Update() error {
	g.updateBackground()
	g.updatePlayer()
	g.updateTurrets()
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	g.renderBackground(screen)
	g.renderPlayer(screen)
	g.renderTurrets(screen)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return g.W, g.H
}

func (g *Game) updatePlayer() {
	if ebiten.IsKeyPressed(ebiten.KeyD) {
		g.pl.pos[0] += 1
	}

	if ebiten.IsKeyPressed(ebiten.KeyA) {
		g.pl.pos[0] -= 1
	}

	if ebiten.IsKeyPressed(ebiten.KeyW) {
		g.pl.pos[1] -= 1
	}

	if ebiten.IsKeyPressed(ebiten.KeyS) {
		g.pl.pos[1] += 1
	}
}

func (g *Game) updateBackground() {
	g.bg.pos[0] -= 2
	g.bg.pos[0] = float64(int(g.bg.pos[0]) % g.W)
}

func (g *Game) updateTurrets() {
	x, y := ebiten.CursorPosition()

	for t, offset := range g.turrets {

		t.pos[0] = g.pl.pos[0] + offset[0]
		t.pos[1] = g.pl.pos[1] + offset[1]

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

func (g *Game) renderBackground(screen *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(g.bg.pos[0], g.bg.pos[1])
	screen.DrawImage(g.bg.img, op)

	op = &ebiten.DrawImageOptions{}
	op.GeoM.Translate(g.bg.pos[0]+float64(g.W), g.bg.pos[1])
	screen.DrawImage(g.bg.img, op)
}

func (g *Game) renderPlayer(screen *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(g.pl.pos[0]-float64(g.pl.w)/2, g.pl.pos[1]-float64(g.pl.h)/2)
	op.GeoM.Scale(2, 2)
	screen.DrawImage(g.pl.img, op)
}

func (g *Game) renderTurrets(screen *ebiten.Image) {
	for t, _ := range g.turrets {
		w, h := t.img.Size()

		op := &ebiten.DrawImageOptions{}

		op.GeoM.Translate(-float64(w)/2, -float64(h)/2)
		op.GeoM.Rotate(float64(t.rotation%360) * 2 * math.Pi / 360)

		op.GeoM.Translate(t.pos[0]*2, t.pos[1]*2)
		op.GeoM.Scale(1, 1)
		screen.DrawImage(t.img, op)
	}
}
