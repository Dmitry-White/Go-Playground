package main

import "sort"

func median(nums []float64) float64 {
	// Sorting mutates, making a copy to avoid mutating original slice
	vals := make([]float64, len(nums))
	copy(vals, nums)

	sort.Float64s(vals)

	i := len(vals) / 2
	if len(vals)%2 == 1 {
		// Case: odd number of values
		return vals[i]
	}
	// Case: even number of values
	return (vals[i-1] + vals[i]) / 2
}
