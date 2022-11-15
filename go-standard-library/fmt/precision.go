package main

import "fmt"

func precision() {
	f := 123.4567

	// TODO: Control the decimal precision
	fmt.Printf("%.2f\n", f)

	// TODO: Print with width 10 and default precision
	fmt.Printf("%10f\n", f)

	// TODO: Print with padding and precision
	fmt.Printf("%10.2f\n", f)

	// TODO: Always use a sign, position or negative
	fmt.Printf("%+10.2f\n", f)

	// TODO: Pad with 0s instead of spaces
	fmt.Printf("%010.2f\n", f)
}
