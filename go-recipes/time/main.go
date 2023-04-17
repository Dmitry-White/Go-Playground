package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("------------ NextDay -------------")
	// Date: 2021-12-31 00:00:00 +0000 UTC Friday
	// Next: 2022-01-03 00:00:00 +0000 UTC Monday
	fmt.Println(nextDay(time.Date(2021, time.December, 31, 0, 0, 0, 0, time.UTC)))

	// Date: 2022-01-04 00:00:00 +0000 UTC Tuesday
	// Next: 2022-01-05 00:00:00 +0000 UTC Wednesday
	fmt.Println(nextDay(time.Date(2022, time.January, 4, 0, 0, 0, 0, time.UTC)))
	fmt.Println("----------------------------------")

	fmt.Println("----------- Dot Product ----------")
	fmt.Println(dotProduct(VECTOR, VECTOR))
	fmt.Println(withTimeMeasure(dotProduct, VECTOR, VECTOR))
	fmt.Println("----------------------------------")

	fmt.Println("------------- Birthday -----------")
	lenon := time.Date(1940, time.October, 9, 18, 30, 0, 0, time.UTC)
	fmt.Println(getBirthday(lenon)) // 3.5s
	fmt.Println("----------------------------------")
}
