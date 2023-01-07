package game

import (
	"LudumDare/assets"
	"log"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/audio"
	"github.com/hajimehoshi/ebiten/v2/audio/mp3"
)

type leveller interface {
	Update() error
	Draw(screen *ebiten.Image)
}

type Game struct {
	bg *entity
	pl *entity
	W  int
	H  int

	currentLevel leveller
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

	g := &Game{
		W:            sW,
		H:            sH,
		currentLevel: NewSplashScreen(sW, sH),
	}

	go func() {
		time.Sleep(2 * time.Second)
		g.currentLevel = NewMainScreen(sW, sH)
	}()

	return g
}

func (g *Game) Update() error {
	return g.currentLevel.Update()
}

func (g *Game) Draw(screen *ebiten.Image) {
	g.currentLevel.Draw(screen)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return g.W, g.H
}
