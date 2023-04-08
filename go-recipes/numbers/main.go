package main

import "log"

func main() {
	log.Println(meanNaive([]int{1, 2, 3}))    // 2
	log.Println(meanNaive([]int{1, 2, 3, 4})) // 2

	log.Println(meanSafe([]int{1, 2, 3}))    // 2
	log.Println(meanSafe([]int{1, 2, 3, 4})) // 2.5
}
