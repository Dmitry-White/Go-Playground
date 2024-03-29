package misc

import "image/color"

type Shape struct {
	x     int
	y     int
	color color.Color
}

type Device interface {
	Set(int, int, color.Color)
}

type Drawer interface {
	Draw(d Device)
}
