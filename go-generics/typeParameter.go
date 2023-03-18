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

// Print prints the information for a solar product.
// The string is enriched with the required kineteco legal information.
func (s *Solar) Print() string {
	return fmt.Sprintf("%s - %v\n", kinetecoPrint, *s)
}

// Print prints the information for a wind product.
// The string is enriched with the required kineteco legal information.
func (w *Wind) Print() string {
	return fmt.Sprintf("%s - %v\n", kinetecoPrint, *w)
}

var kinetecoPrint string = "Kineteco Deal:"

func typeParameter() {
	// Create three different energy offers of kineteco
	solar2k := Solar{Name: "Solar 2000", Netto: 4.500}
	solar3k := Solar{Name: "Solar 3000", Netto: 4.000}
	windwest := Wind{Name: "Wind West", Netto: 3.950}

	// Print details for each energy offer with kineteco branding
	fmt.Println(solar2k.Print())
	fmt.Println(solar3k.Print())
	fmt.Println(windwest.Print())
}
