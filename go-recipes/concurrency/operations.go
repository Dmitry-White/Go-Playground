package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"net/http"
	"strconv"
	"time"
)

type Result struct {
	date time.Time
	dist float64
	err  error
}

const (
	DATE_FORMAT = "2006-01-02"
	TIME_FORMAT = "2006-01-02T15:04:05"
	BASE_URL    = "http://localhost:8080"
)

// ############################## HTTP Server ######################################
func csvHandler(w http.ResponseWriter, r *http.Request) {
	s := r.URL.Path[1:] // trim leading /

	day, err := time.Parse(DATE_FORMAT, s)
	if err != nil {
		msg := fmt.Sprintf("bad date: %q", s)
		log.Printf("error: %s", msg)
		http.Error(w, msg, http.StatusBadRequest)
		return
	}

	fmt.Fprintf(w, "start,end,distance,passengers\n")

	start := day
	nextDay := day.Add(24 * time.Hour)
	for start.Before(nextDay) {
		start = start.Add(3 * time.Second)
		end := start.Add(17 * time.Second)
		distance := 3.14
		passengers := 1
		fmt.Fprintf(w, "%s,%s,%.2f,%d\n", start.Format(TIME_FORMAT), end.Format(TIME_FORMAT), distance, passengers)
	}
}

func csvServer() {
	router := http.ServeMux{}
	router.HandleFunc("/", csvHandler)

	err := http.ListenAndServe(OPERATIONS_SERVER_PORT, &router)
	if err != nil {
		log.Fatal(err)
	}
}

// #################################################################################

// ############################## HTTP Client ######################################
func dayDistance(r io.Reader) (float64, error) {
	rdr := csv.NewReader(r)
	total, lNum := 0.0, 0
	for {
		//2021-01-02T23:58:36,2021-01-02T23:58:40,3.41,1
		fields, err := rdr.Read()
		if err == io.EOF {
			break
		}

		if err != nil {
			return 0, err
		}

		lNum++
		if lNum == 1 {
			continue // skip header
		}

		dist, err := strconv.ParseFloat(fields[2], 64)
		if err != nil {
			return 0, err
		}

		total += dist
	}

	return total, nil
}

func monthDistanceSequential(month time.Time) (float64, error) {
	totalDistance := 0.0

	date := month
	for date.Month() == month.Month() {
		url := fmt.Sprintf("%s/%s", BASE_URL, date.Format(DATE_FORMAT))
		resp, err := http.Get(url)
		if err != nil {
			return 0, err
		}

		if resp.StatusCode != http.StatusOK {
			return 0, fmt.Errorf("bad status: %d %s", resp.Request.Response.StatusCode, resp.Status)
		}

		defer resp.Body.Close()
		dist, err := dayDistance(resp.Body)
		if err != nil {
			return 0, err
		}
		totalDistance += dist
		date = date.Add(24 * time.Hour)
	}

	return totalDistance, nil
}

func dateWorker(date time.Time, ch chan<- Result) {
	res := Result{date: date}
	defer func() {
		ch <- res
	}()

	url := fmt.Sprintf("%s/%s", BASE_URL, date.Format(DATE_FORMAT))
	resp, err := http.Get(url)
	if err != nil {
		res.err = err
		return
	}

	if resp.StatusCode != http.StatusOK {
		res.err = fmt.Errorf("bad status: %d %s", resp.Request.Response.StatusCode, resp.Status)
		return
	}

	defer resp.Body.Close()
	res.dist, res.err = dayDistance(resp.Body)
}

func monthDistanceConcurrent(month time.Time) (float64, error) {
	totalDistance := 0.0
	numWorkers := 0
	ch := make(chan Result)

	date := month
	for date.Month() == month.Month() {
		go dateWorker(date, ch)
		numWorkers++
		date = date.Add(24 * time.Hour)
	}

	for i := 0; i < numWorkers; i++ {
		res := <-ch
		if res.err != nil {
			return 0, res.err
		}
		totalDistance += res.dist
	}

	return totalDistance, nil
}

func getDistance(strategy string) func(month time.Time) interface{} {
	operationsFunc, err := getOperationsFunc(strategy)
	if err != nil {
		log.Fatalln(err)
	}

	return func(month time.Time) interface{} {
		start := time.Now()

		distance, err := operationsFunc(month)
		if err != nil {
			log.Fatal(err)
		}

		duration := time.Since(start)

		result := fmt.Sprintf("distance=%.2f, duration=%v\n", distance, duration)

		return result
	}
}

// #################################################################################
