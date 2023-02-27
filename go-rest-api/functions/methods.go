package main

type Dimension struct {
	length, width, height int
}

func (d Dimension) Area() int {
	return d.length * d.width
}

func (d Dimension) Volume() int {
	return d.length * d.width * d.height
}
