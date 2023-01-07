package game

import (
	"LudumDare/assets"
	"LudumDare/entities"

	"github.com/hajimehoshi/ebiten/v2"
)

type SplashScreen struct {
	w  int
	h  int
	bg *entities.Entity
}

func NewSplashScreen(sW, sH int) *SplashScreen {
	return &SplashScreen{
		w: sW,
		h: sH,
		bg: &entities.Entity{
			Img: ebiten.NewImageFromImage(assets.Splash()),
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
	screen.DrawImage(s.bg.Img, op)
}
