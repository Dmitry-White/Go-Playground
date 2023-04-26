package misc

import "image/color"

type Rectangle struct {
	Shape
	h int
	w int
}

func NewRectangle(x, y, h, w int, c color.Color) *Rectangle {
	rectangle := Rectangle{Shape{x, y, c}, h, w}
	return &rectangle
}

func (r *Rectangle) Draw(d Device) {
	minX := r.x - r.h/2
	minY := r.y - r.w/2
	maxX := r.x + r.h/2
	maxY := r.y + r.w/2

	for x := minX; x <= maxX; x++ {
		for y := minY; y <= maxY; y++ {
			d.Set(x, y, r.c)
		}
	}
}
