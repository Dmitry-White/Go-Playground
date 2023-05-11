package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"sync/atomic"

	"golang.org/x/sync/errgroup"
)

/*
	More Info: https://www.fullstory.com/blog/why-errgroup-withcontext-in-golang-server-handlers/
*/

const (
	URL_TEMPLATE = "https://d37ci6vzurychx.cloudfront.net/trip-data/%s_tripdata_2020-%02d.parquet"
	START_MONTH  = 1
	END_MONTH    = 12
)

func downloadSize(url string) (int, error) {
	resp, err := http.Head(url)
	if err != nil {
		return 0, err
	}

	if resp.StatusCode != http.StatusOK {
		return 0, fmt.Errorf("bad status: %d %s", resp.StatusCode, resp.Status)
	}

	return strconv.Atoi(resp.Header.Get("Content-Length"))
}

func downloadHanlder(dlSize *int64, url string) func() error {
	return func() error {
		log.Print("URL: ", url)

		size, err := downloadSize(url)
		if err != nil {
			return err
		}

		atomic.AddInt64(dlSize, int64(size))
		return nil
	}
}

func analyzeURLs(urls []string) (int, error) {
	dlSize := int64(0)
	eg := errgroup.Group{}

	for _, url := range urls {
		eg.Go(downloadHanlder(&dlSize, url))
	}

	err := eg.Wait()
	if err != nil {
		return 0, err
	}

	return int(dlSize), nil
}

func generateURLs() []string {
	urls := []string{}
	vendors := []string{"yellow", "green"}

	for _, vendor := range vendors {
		for month := START_MONTH; month <= END_MONTH; month++ {
			url := fmt.Sprintf(URL_TEMPLATE, vendor, month)
			urls = append(urls, url)
		}
	}

	return urls
}

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
