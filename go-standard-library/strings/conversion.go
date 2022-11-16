package main

import (
	"fmt"
	"strconv"
)

// The strconv package contains a variety of functions
// for parsing and converting numbers, values and strings
func conversion() {
	sampleInt := 100
	sampleStr := "250"

	// TODO: Perform character coversion, not a number one
	newstr := string(sampleInt)
	fmt.Println("Result of using string(): ", newstr)

	// TODO: Convert an integer to string
	s1 := strconv.Itoa(sampleInt)
	fmt.Printf("%T, %v\n", s1, s1)

	// TODO: Convert a string to integer
	i2, err := strconv.Atoi(sampleStr)
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Printf("%T, %v\n", i2, i2)

	// TODO: Other parse functions
	b, _ := strconv.ParseBool("true")
	fmt.Println(b)
	f, _ := strconv.ParseFloat("3.14159", 64)
	fmt.Println(f)
	i, _ := strconv.ParseInt("-42", 10, 64)
	fmt.Println(i)
	u, _ := strconv.ParseUint("42", 10, 64)
	fmt.Println(u)

	// TODO: Other format functions
	fb := strconv.FormatBool(true)
	fmt.Println(fb)
	ff := strconv.FormatFloat(3.14159, 'E', -1, 64)
	fmt.Println(ff)
	fi := strconv.FormatInt(-42, 10)
	fmt.Println(fi)
	fu := strconv.FormatUint(42, 10)
	fmt.Println(fu)
}
