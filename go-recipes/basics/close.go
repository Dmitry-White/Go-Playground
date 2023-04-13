package main

import (
	"fmt"
)

type Item struct {
	SKU  string
	Name string
}

func writeItems(filename string, items []Item) error {
	fmt.Println("Not Implemented")
	fmt.Printf("Input %v into %s\n", items, filename)
	return nil
}

func close(filename string) error {
	items := []Item{
		{"m183x", "Magic Wand"},
		{"m184y", "Invisibility Cape"},
		{"m185z", "Levitation Spell"},
	}

	err := writeItems(filename, items)
	if err != nil {
		return err
	}
	return nil
}
