package parameter

import "fmt"

func TypeParameter() {
	// Create three different energy offers of kineteco
	solar2k := Solar{Name: "Solar 2000", Netto: 4.500}
	solar3k := Solar{Name: "Solar 3000", Netto: 4.000}
	windwest := Wind{Name: "Wind West", Netto: 3.950}

	// Print details for each energy offer with kineteco branding
	fmt.Println(solar2k.print())
	fmt.Println(solar3k.print())
	fmt.Println(windwest.print())
	fmt.Println(printGeneric[Solar](solar2k))

	solarSlice := []Solar{solar2k, solar3k}
	windSlice := []Wind{windwest, windwest}
	printSlice[Solar](solarSlice)
	printSlice[Wind](windSlice)
}
