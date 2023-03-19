package inference

import "golang.org/x/exp/constraints"

// Solar handles all the different energy offers powered by solar.
type Solar struct {
	Name  string
	Netto float64
}

// Wind handles all the different energy offers powered by wind.
type Wind struct {
	Name  string
	Netto float64
}

type Energy interface {
	Solar | Wind
	Cost() float64
}

// Number is either a floating point number or an integer
type Number interface {
	constraints.Float | constraints.Integer
}
