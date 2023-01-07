package assets

import (
	"bytes"
	"image"
	"image/png"
	"log"

	_ "embed"
)

//go:embed bg.png
var bg []byte

func BG() image.Image {
	img, err := png.Decode(bytes.NewReader(bg))
	if err != nil {
		log.Fatal(err)
	}
	return img
}

//go:embed pl.png
var pl []byte

func PL() image.Image {
	img, err := png.Decode(bytes.NewReader(pl))
	if err != nil {
		log.Fatal(err)
	}
	return img
}
