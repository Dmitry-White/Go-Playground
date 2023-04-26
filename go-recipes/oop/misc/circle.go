package misc

import (
	"image/color"
	"math"
)

type Circle struct {
	Shape
	radius int
}

func NewCircle(x, y, radius int, color color.Color) *Circle {
	circle := Circle{Shape{x, y, color}, radius}
	return &circle
}

func (c *Circle) Draw(d Device) {
	minX := c.x - c.radius
	minY := c.y - c.radius
	maxX := c.x + c.radius
	maxY := c.y + c.radius

	for x := minX; x <= maxX; x++ {
		for y := minY; y <= maxY; y++ {
			dx := x - c.x
			dy := y - c.y
			point := int(math.Sqrt(float64(dx*dx + dy*dy)))

			if point <= c.radius {
				d.Set(x, y, c.color)
			}
		}
	}
}
