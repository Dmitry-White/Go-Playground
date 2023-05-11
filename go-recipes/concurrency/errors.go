package main

import (
	"fmt"
)

/*
	More Info: https://www.fullstory.com/blog/why-errgroup-withcontext-in-golang-server-handlers/
*/

const (
	URL_TEMPLATE = "https://d37ci6vzurychx.cloudfront.net/trip-data/%s_tripdata_2020-%02d.parquet"
	START_MONTH  = 1
	END_MONTH    = 12
)

func getErrors() interface{} {
	urls := generateURLs()

	size, err := analyzeURLs(urls)
	if err != nil {
		return err
	}

	sizeGB := float64(size) / (1 << 30)
	result := fmt.Sprintf("size = %.2fGB\n", sizeGB)

	return result
}
