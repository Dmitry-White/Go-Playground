package misc

import (
	"fmt"
	"image"
	"image/png"
	"io"
)

type Canvas struct {
	width  int
	height int
	shapes []Drawer
}

func NewCanvas(width, height int) (*Canvas, error) {
	if width <= 0 || height <= 0 {
		err := fmt.Errorf("Negative size: width=%d, height=%d", width, height)
		return nil, err
	}

	canvas := Canvas{width, height, []Drawer{}}
	return &canvas, nil
}

func (ic *Canvas) Add(d Drawer) {
	ic.shapes = append(ic.shapes, d)
}

func (ic *Canvas) Draw(w io.Writer) error {
	canvas := image.Rect(0, 0, ic.width, ic.height)
	image := image.NewRGBA(canvas)

	for _, shape := range ic.shapes {
		shape.Draw(image)
	}

	return png.Encode(w, image)
}
