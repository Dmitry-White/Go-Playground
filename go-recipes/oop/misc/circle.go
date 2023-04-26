package misc

import (
	"image/color"
	"math"
)

type Circle struct {
	Shape
	r int
}

func NewCircle(x, y, r int, c color.Color) *Circle {
	circle := Circle{Shape{x, y, c}, r}
	return &circle
}

func (c *Circle) Draw(d Device) {
	minX := c.x - c.r
	minY := c.y - c.r
	maxX := c.x + c.r
	maxY := c.y + c.r

	for x := minX; x <= maxX; x++ {
		for y := minY; y <= maxY; y++ {
			dx := x - c.x
			dy := y - c.y
			point := int(math.Sqrt(float64(dx*dx + dy*dy)))

			if point <= c.r {
				d.Set(x, y, c.c)
			}
		}
	}
}
