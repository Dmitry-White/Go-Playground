package main

import "fmt"

func dimensions(length, width, height int) (area int, volume int) {
	area = length * width
	volume = area * height

	return
}

func main() {
	area, volume := dimensions(5, 5, 5)
	fmt.Printf("area: %d, volume: %d\n", area, volume)

	dimension := Dimension{5, 5, 5}
	fmt.Printf("area: %d, volume: %d\n", dimension.Area(), dimension.Volume())
}
