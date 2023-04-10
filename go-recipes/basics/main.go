package main

import (
	"fmt"
)

func main() {
	fmt.Println("-------------- Mean --------------")
	fmt.Println(meanNaive([]int{1, 2, 3}))    // 2
	fmt.Println(meanNaive([]int{1, 2, 3, 4})) // 2

	fmt.Println(meanSafe([]int{1, 2, 3}))    // 2
	fmt.Println(meanSafe([]int{1, 2, 3, 4})) // 2.5
	fmt.Println("----------------------------------")

	fmt.Println("------------- Median -------------")
	fmt.Println(median([]float64{3, 1, 2}))    // 2
	fmt.Println(median([]float64{3, 1, 4, 2})) // 2.5
	fmt.Println("----------------------------------")

	fmt.Println("----------- Frequency ------------")
	fmt.Println(frequency([]string{}))
	commaOk() // The price of Banana is $0. We don't have Apples.
	fmt.Println("----------------------------------")
}
