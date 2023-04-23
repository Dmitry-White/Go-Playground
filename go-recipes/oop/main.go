package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("------------- Struct -------------")
	fmt.Printf("%+v\n", getDoorEvent("front door", time.Now(), "open"))
	fmt.Printf("%+v\n", getTermostatEvent("front door", time.Now(), 25))
	fmt.Println("----------------------------------")
}
