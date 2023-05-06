package main

import "fmt"

func main() {
	fmt.Println("------------- Client -------------")
	fmt.Printf("%+v\n", sendMetrics())
	fmt.Println("----------------------------------")

	fmt.Println("-------------- Auth --------------")
	fmt.Printf("%+v\n", authenticate())
	fmt.Println("----------------------------------")

	fmt.Println("------------- Server -------------")
	fmt.Printf("%+v\n", startServer())
	fmt.Println("----------------------------------")
}
