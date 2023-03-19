package inference

import "fmt"

func TypeInference() {
	// Create three different energy offers of kineteco
	solar2k := Solar{Name: "Solar 2000", Netto: 4.500}
	solar3k := Solar{Name: "Solar 3000", Netto: 4.000}
	windwest := Wind{Name: "Wind West", Netto: 3.950}

	// Print details for each energy offer with kineteco branding
	fmt.Println(solar2k.print())
	fmt.Println(solar3k.print())
	fmt.Println(windwest.print())
	fmt.Println(printGeneric(solar2k))

	solarSlice := []Solar{solar2k, solar3k}
	windSlice := []Wind{windwest, windwest}
	printSlice(solarSlice)
	printSlice(windSlice)

	Cost(10, solar2k.Netto)
	// Cost(0.45, 10) // Compile error since 10 is not infered as float64
	Cost(0.45, float64(10)) // Either cast explicitly
	Cost[float64](0.45, 10) // Or add type parameter

	newSolarSlice := SolarSlice{solar2k, solar3k}
	printSlice(newSolarSlice)            // Incorrectly infers the constrained type of the input slice
	printSliceConstrained(newSolarSlice) // Correctly infers the constrained type of the input slice
}
