package main

func doDefer() error {
	nums := []int{1}
	panickingHandler(nums)
	return nil
}

func panickingHandler(nums []int) int {
	return nums[len(nums)-2]
}
