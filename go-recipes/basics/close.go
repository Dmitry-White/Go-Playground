package main

import (
	"encoding/csv"
	"os"
)

type Item struct {
	SKU  string
	Name string
}

func writeItems(filename string, items []Item) error {
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	header := []string{"sku", "name"}

	csvWriter := csv.NewWriter(file)
	defer csvWriter.Flush()

	err = csvWriter.Write(header)
	if err != nil {
		return err
	}

	for _, item := range items {
		row := []string{item.SKU, item.Name}

		err := csvWriter.Write(row)
		if err != nil {
			return err
		}
	}

	return csvWriter.Error()
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
