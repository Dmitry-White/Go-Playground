package main

import "fmt"

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

func PrintGeneric[T any](s T) string {
	return fmt.Sprintf("%s - %v\n", kinetecoPrint, s)
}

func typeParameter() {
	// Create three different energy offers of kineteco
	solar2k := Solar{Name: "Solar 2000", Netto: 4.500}
	solar3k := Solar{Name: "Solar 3000", Netto: 4.000}
	windwest := Wind{Name: "Wind West", Netto: 3.950}

	// Print details for each energy offer with kineteco branding
	fmt.Println(solar2k.print())
	fmt.Println(solar3k.print())
	fmt.Println(windwest.print())
	fmt.Println(PrintGeneric[Solar](solar2k))
}
