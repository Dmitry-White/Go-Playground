package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("------------- Struct -------------")
	fmt.Printf("%+v\n", getDoorEvent("front door", time.Now(), "open"))
	fmt.Println("----------------------------------")
}
