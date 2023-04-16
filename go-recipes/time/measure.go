package main

import (
	"fmt"
	"log"
	"reflect"
	"runtime"
	"time"
)

var VECTOR = []float64{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

/*
Variadic Function
https://gobyexample.com/variadic-functions
*/
type OriginalFunc = func(...[]float64) (float64, error)

func timeit(name string) func() {
	start := time.Now()

	return func() {
		duration := time.Since(start)
		log.Printf("%s took %s", name, duration)
	}
}

func withTimeMeasure(original OriginalFunc, args ...[]float64) (float64, error) {
	name := runtime.FuncForPC(reflect.ValueOf(original).Pointer()).Name()
	defer timeit(name)()

	return original(args...)
}

/*
Calculates a Dot Product that returns a scalar value representing the amount of "force" applied.
https://www.mathsisfun.com/algebra/vectors-dot-product.html
https://math.libretexts.org/Bookshelves/Calculus/Calculus_(OpenStax)/12%3A_Vectors_in_Space/12.03%3A_The_Dot_Product
*/
func dotProduct(vs ...[]float64) (float64, error) {
	v1 := vs[0]
	v2 := vs[1]

	if len(v1) != len(v2) {
		return 0, fmt.Errorf("dot of different size vectors")
	}

	time.Sleep(time.Second)

	d := 0.0
	for i, val1 := range v1 {
		val2 := v2[i]
		d += val1 * val2
	}

	return d, nil
}
