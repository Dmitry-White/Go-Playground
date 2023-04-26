package misc

import (
	"image/color"
	"log"
	"os"
)

const IMAGE_NAME = "face.png"

var (
	Red   = color.RGBA{0xFF, 0, 0, 0xFF}
	Green = color.RGBA{0, 0xFF, 0, 0xFF}
	Blue  = color.RGBA{0, 0, 0xFF, 0xFF}
)

func GetFace() *Canvas {
	ic, err := NewCanvas(200, 200)
	if err != nil {
		log.Fatal(err)
	}

	ic.Add(NewCircle(100, 100, 80, Green))
	ic.Add(NewCircle(60, 60, 10, Blue))
	ic.Add(NewCircle(140, 60, 10, Blue))
	ic.Add(NewRectangle(100, 130, 80, 10, Red))

	f, err := os.Create(IMAGE_NAME)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	err = ic.Draw(f)
	if err != nil {
		log.Fatal(err)
	}

	return ic
}
