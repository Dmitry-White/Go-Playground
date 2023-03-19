package parameter

import "fmt"

var kinetecoPrint string = "Kineteco Deal:"

// Print prints the information for a solar product.
// The string is enriched with the required kineteco legal information.
func (s *Solar) print() string {
	return fmt.Sprintf("%s - %v\n", kinetecoPrint, *s)
}

// Print prints the information for a wind product.
// The string is enriched with the required kineteco legal information.
func (w *Wind) print() string {
	return fmt.Sprintf("%s - %v\n", kinetecoPrint, *w)
}

func printGeneric[T any](s T) string {
	return fmt.Sprintf("%s - %v\n", kinetecoPrint, s)
}

func printSlice[T any](sl []T) {
	for i, el := range sl {
		fmt.Printf("%d: %s", i, printGeneric[T](el))
	}
}
