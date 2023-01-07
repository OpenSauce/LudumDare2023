package game

import (
	"LudumDare/assets"

	"github.com/hajimehoshi/ebiten/v2"
)

type SplashScreen struct {
	w  int
	h  int
	bg *entity
}

func NewSplashScreen(sW, sH int) *SplashScreen {
	return &SplashScreen{
		w: sW,
		h: sH,
		bg: &entity{
			img: ebiten.NewImageFromImage(assets.Splash()),
		},
	}
}

func (s *SplashScreen) Update() error {
	return nil
}

func (s *SplashScreen) Draw(screen *ebiten.Image) {
	s.renderBackground(screen)
}

func (s *SplashScreen) renderBackground(screen *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}
	screen.DrawImage(s.bg.img, op)
}
