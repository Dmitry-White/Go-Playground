package main

func sum(nums []int) int {
	total := 0
	for _, n := range nums {
		total += n
	}

	return total
}

func meanNaive(nums []int) int {
	s := sum(nums)
	return s / len(nums)
}

func meanSafe(nums []int) float64 {
	s := sum(nums)
	return float64(s) / float64(len(nums))
}
