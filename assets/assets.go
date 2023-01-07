package assets

import (
	"bytes"
	"image"
	"image/png"
	"io"
	"log"
	"os"

	_ "embed"
)

//go:embed background.png
var bg []byte

func Background() image.Image {
	img, err := png.Decode(bytes.NewReader(bg))
	if err != nil {
		log.Fatal(err)
	}
	return img
}

//go:embed turtle.png
var pl []byte

func Turtle() image.Image {
	img, err := png.Decode(bytes.NewReader(pl))
	if err != nil {
		log.Fatal(err)
	}
	return img
}

//go:embed turret.png
var turret []byte

func Turret() image.Image {
	img, err := png.Decode(bytes.NewReader(turret))
	if err != nil {
		log.Fatal(err)
	}
	return img
}

func LoadMusic() io.Reader {
	f, err := os.Open("assets/Doomsayer.mp3")
	if err != nil {
		log.Fatal("Error loading music")
	}
	return f
}
