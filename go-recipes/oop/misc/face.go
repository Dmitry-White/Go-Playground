package misc

import (
	"image/color"
	"io"
	"log"
	"os"
)

const IMAGE_NAME = "face.png"

var (
	Red   = color.RGBA{0xFF, 0, 0, 0xFF}
	Green = color.RGBA{0, 0xFF, 0, 0xFF}
	Blue  = color.RGBA{0, 0, 0xFF, 0xFF}
)

type Circle struct {
	// FIXME
}

func NewCircle(x, y, r int, c color.Color) *Circle {
	// FIXME
	return nil
}

type Rectangle struct {
	// FIXME
}

func NewRectangle(x, y, h, w int, c color.Color) *Rectangle {
	// FIXME
	return nil
}

type Device interface {
	Set(int, int, color.Color)
}

type ImageCanvas struct {
	// FIXME
}

func NewImageCanvas(width, height int) (*ImageCanvas, error) {
	// FIXME
	return nil, nil
}

type Drawer interface {
	// FIXME
}

func (ic *ImageCanvas) Add(d Drawer) {
	// FIXME
}

func (ic *ImageCanvas) Draw(w io.Writer) error {
	// FIXME
	return nil
}

func GetFace() *ImageCanvas {
	ic, err := NewImageCanvas(200, 200)
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
	if err := ic.Draw(f); err != nil {
		log.Fatal(err)
	}

	return ic
}
