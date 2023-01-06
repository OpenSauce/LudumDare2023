package assets

import (
	"bytes"
	"image"
	"image/png"
	"log"

	_ "embed"
)

//go:embed background/starwars.png
var bg []byte

func Background() image.Image {
	img, err := png.Decode(bytes.NewReader(bg))
	if err != nil {
		log.Fatal(err)
	}
	return img
}
