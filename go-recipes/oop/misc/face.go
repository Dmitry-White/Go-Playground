package misc

import (
	"fmt"
	"image"
	"image/color"
	"image/png"
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

func NewCircle(x, y, r int, c color.Color) *Circle {
	circle := Circle{Shape{x, y, c}, r}
	return &circle
}

func NewRectangle(x, y, h, w int, c color.Color) *Rectangle {
	rectangle := Rectangle{Shape{x, y, c}, h, w}
	return &rectangle
}

func NewImageCanvas(width, height int) (*ImageCanvas, error) {
	imageCanvas := ImageCanvas{width, height, []Drawer{}}
	return &imageCanvas, nil
}

func (s *Shape) Draw(d Device) {
	d.Set(s.x, s.y, s.c)
}

func (ic *ImageCanvas) Add(d Drawer) {
	ic.shapes = append(ic.shapes, d)
}

func (ic *ImageCanvas) Draw(w io.Writer) error {
	canvas := image.Rect(0, 0, ic.width, ic.height)
	image := image.NewRGBA(canvas)

	fmt.Printf("%#v", ic.shapes)

	for _, shape := range ic.shapes {
		shape.Draw(image)
	}

	return png.Encode(w, image)
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
