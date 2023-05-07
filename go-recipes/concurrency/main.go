package main

import "fmt"

func main() {
	fmt.Println("----------- Operations -----------")
	fmt.Printf("%+v\n", sequential())
	fmt.Printf("%+v\n", concurrent())
	fmt.Println("----------------------------------")
}
