package main

import (
	"fmt"
	"log"
	"runtime"
	"sort"
	"sync"
	"time"
)

/*
	More Info: https://gobyexample.com/worker-pools
*/

func calcMedian(values []float64) float64 {
	nums := make([]float64, len(values))

	copy(nums, values)
	sort.Float64s(nums)

	i := len(nums) / 2
	if len(nums)%2 == 1 {
		return nums[i]
	}

	return (nums[i-1] + nums[i]) / 2
}

func worker(todoJobs <-chan []float64, wg *sync.WaitGroup, doneJobs chan<- float64) {
	for values := range todoJobs {
		median := calcMedian(values)
		log.Printf("Calculation Done %v -> %f", values, median)

		doneJobs <- median
		wg.Done()
	}

	log.Printf("Shutting down...")
}

func prepareWorkers(todoJobs <-chan []float64, wg *sync.WaitGroup, doneJobs chan<- float64) {
	log.Printf("Preparing workers...")

	workerPoolSize := runtime.NumCPU()

	for i := 0; i < workerPoolSize; i++ {
		go worker(todoJobs, wg, doneJobs)
	}
	log.Printf("Worker Pool: %d\n", workerPoolSize)

	log.Printf("Workers prepared!")
}

func submitJobs(todoJobs chan<- []float64, wg *sync.WaitGroup, vectors [][]float64) {
	log.Printf("Submitting jobs...")

	jobPoolSize := len(vectors)

	for _, vec := range vectors {
		wg.Add(1)
		todoJobs <- vec
	}
	log.Printf("Job Pool: %d\n", jobPoolSize)

	log.Printf("Jobs submitted!")
}

func gatherResults(doneJobs <-chan float64) []float64 {
	log.Printf("Gathering results...")

	results := []float64{}

	for median := range doneJobs {
		log.Println("Median: ", median)
		results = append(results, median)
	}

	log.Printf("Results gathered!")
	return results
}

func calcMultiMedian(vectors [][]float64) []float64 {
	jobPoolSize := len(vectors)
	todoJobs := make(chan []float64, jobPoolSize)
	doneJobs := make(chan float64, jobPoolSize)

	wg := sync.WaitGroup{}

	prepareWorkers(todoJobs, &wg, doneJobs)
	submitJobs(todoJobs, &wg, vectors)

	wg.Wait()
	close(todoJobs)
	close(doneJobs)

	results := gatherResults(doneJobs)

	return results
}

func multiMedian() []float64 {
	vectors := [][]float64{
		{1.1, 2.2, 3.3},
		{2.2, 3.3, 4.4},
		{3.3, 4.4, 5.5},
		{4.4, 5.5, 6.6},
		{5.5, 6.6, 7.7},
	}

	multiMedian := calcMultiMedian(vectors)
	time.Sleep(100 * time.Millisecond) // Let workers terminate
	fmt.Println("DONE")

	return multiMedian
}
