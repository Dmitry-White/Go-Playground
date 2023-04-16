package main

import "fmt"

var VECTOR = []float64{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

/*
	Calculates a Dot Product that returns a scalar value representing the amount of "force" applied.
	https://www.mathsisfun.com/algebra/vectors-dot-product.html
	https://math.libretexts.org/Bookshelves/Calculus/Calculus_(OpenStax)/12%3A_Vectors_in_Space/12.03%3A_The_Dot_Product
*/
func dotProduct(v1, v2 []float64) (float64, error) {
	if len(v1) != len(v2) {
		return 0, fmt.Errorf("dot of different size vectors")
	}

	d := 0.0
	for i, val1 := range v1 {
		val2 := v2[i]
		d += val1 * val2
	}

	return d, nil
}
