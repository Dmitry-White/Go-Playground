package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("------------- Struct -------------")
	fmt.Println(getDoorEvent("front door", time.Now(), "open"))
	fmt.Println("----------------------------------")
}
