package main

import (
	"fmt"
	"math"
)

func basic() {
	// TODO: Print Pi
	fmt.Println(math.Pi)

	// TODO: Ceiling and Floor functions
	fmt.Println(math.Ceil(math.Pi))
	fmt.Println(math.Floor(math.Pi))

	// TODO: Truncate will return int value of X
	fmt.Printf("%.2f\n", math.Trunc(math.Pi))

	// TODO: Max and Min values
	fmt.Println(math.Max(math.Pi, math.Ln2))
	fmt.Println(math.Min(math.Pi, math.Ln2))

	// TODO: Mod operator is like % but for floats
	fmt.Println(17 % 5)
	fmt.Println(math.Mod(math.Pi, math.Ln2))

	// TODO: Round and RoundToEven
	fmt.Printf("%.1f\n", math.Round(10.5))
	fmt.Printf("%.1f\n", math.Round(-10.5))

	fmt.Printf("%.1f\n", math.RoundToEven(10.5))
	fmt.Printf("%.1f\n", math.RoundToEven(11.5))
}
