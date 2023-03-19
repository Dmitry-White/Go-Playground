package main

import (
	"fmt"
	"go-generics/inference"
	"go-generics/parameter"
	"go-generics/set"
)

func main() {
	fmt.Println("----------- Type Parameter -----------")
	parameter.TypeParameter()
	fmt.Println("--------------------------------------")

	fmt.Println("-------------- Type Set --------------")
	set.TypeSet()
	fmt.Println("--------------------------------------")

	fmt.Println("----------- Type Inference -----------")
	inference.TypeInference()
	fmt.Println("--------------------------------------")
}
