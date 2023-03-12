package db

import (
	"encoding/csv"
	"fmt"
	"os"
	"strconv"
	"sync"

	"go-applied-concurrency/api/models"
)

const ProductInputPath string = "./api/db/data/products.csv"
const ProductInputPathTest string = "../db/data/products.csv"

// importProducts imports the start position of the products DB
// Naive approach using built-in Map
func ImportProducts( /*products map[string]models.Product*/ products *sync.Map, dbPath string) error {
	input, err := readCsv(dbPath)
	if err != nil {
		return err
	}

	// each line in the file is a product
	for _, line := range input {
		// bad csv line continue import
		if len(line) != 5 {
			continue
		}
		id := line[0]
		stock, err := strconv.Atoi(line[2])
		// bad csv line continue import
		if err != nil {
			continue
		}
		price, err := strconv.ParseFloat(line[4], 64)
		// bad csv line continue import
		if err != nil {
			continue
		}
		// Naive approach using built-in Map
		// products[id] = models.Product{
		// 	ID:    id,
		// 	Name:  fmt.Sprintf("%s(%s)", line[1], line[3]),
		// 	Stock: stock,
		// 	Price: price,
		// }
		productValue := models.Product{
			ID:    id,
			Name:  fmt.Sprintf("%s(%s)", line[1], line[3]),
			Stock: stock,
			Price: price,
		}
		products.Store(id, productValue)
	}
	return nil
}

// readCsv reads file and converts it to an array of strings
// the format of the csv file is hardcoded so we can take some
// error handling liberties for the sake of brevity
func readCsv(filename string) ([][]string, error) {
	f, err := os.Open(filename)
	if err != nil {
		return [][]string{}, err
	}
	defer f.Close()
	lines, err := csv.NewReader(f).ReadAll()
	if err != nil {
		return [][]string{}, err
	}

	return lines, nil
}

// toOrder attempts to convert an interface to an order
// panics if this is not possible
func toOrder(po interface{}) models.Order {
	order, ok := po.(models.Order)
	if !ok {
		panic(fmt.Errorf("error casting %v to order", po))
	}
	return order
}

// toProduct attempts to convert an interface to an product
// panics if this is not possible
func toProduct(po interface{}) models.Product {
	product, ok := po.(models.Product)
	if !ok {
		panic(fmt.Errorf("error casting %v to product", po))
	}
	return product
}
