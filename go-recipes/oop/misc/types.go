package misc

import "image/color"

type Shape struct {
	x int
	y int
	c color.Color
}

type Circle struct {
	Shape
	r int
}

type Rectangle struct {
	Shape
	h int
	w int
}

type Device interface {
	Set(int, int, color.Color)
}

type ImageCanvas struct {
	width  int
	height int
	shapes []Drawer
}

type Drawer interface {
	Draw(d Device)
}
