package misc

import "image/color"

type Rectangle struct {
	Shape
	height int
	width  int
}

func NewRectangle(x, y, height, width int, color color.Color) *Rectangle {
	rectangle := Rectangle{Shape{x, y, color}, height, width}
	return &rectangle
}

func (r *Rectangle) Draw(d Device) {
	minX := r.x - r.height/2
	minY := r.y - r.width/2
	maxX := r.x + r.height/2
	maxY := r.y + r.width/2

	for x := minX; x <= maxX; x++ {
		for y := minY; y <= maxY; y++ {
			d.Set(x, y, r.color)
		}
	}
}
