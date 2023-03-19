package inference

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

// Cost computes the netto cost for a costumer for Solar and returns the value
func (s Solar) Cost() float64 {
	return s.Netto * 0.4
}

// Cost computes the netto cost for a costumer for Wind and returns the value
func (w Wind) Cost() float64 {
	return w.Netto * 0.3
}

func printGeneric[T Energy](s T) string {
	return fmt.Sprintf("%s - %v\n", kinetecoPrint, s)
}

func printSlice[T Energy](sl []T) {
	fmt.Printf("Type of sl: %T\n", sl)
	for i, el := range sl {
		fmt.Printf("%d: %s", i, printGeneric(el))
	}
}

// Correctly infers the constrained type of the input slice
// since S is constrained by the slice of T and its underlying types
func printSliceConstrained[T Energy, S ~[]T](sl S) {
	fmt.Printf("Type of sl: %T\n", sl)
	for i, el := range sl {
		fmt.Printf("%d: %s", i, printGeneric(el))
	}
}

// Cost multiplies usage with netto to compute the cost.
func Cost[T Number](usage, netto T) T {
	return usage * netto
}
