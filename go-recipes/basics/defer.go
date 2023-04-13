package main

import "fmt"

func panickingHandler(nums []int) int {
	return nums[len(nums)-2]
}

func withRecovery(handler func(args []int) int, args []int) (_ interface{}, err error) {
	defer func() {
		if e := recover(); e != nil { // e is interface{}
			err = fmt.Errorf("%v", e)
		}
	}()

	return handler(args), nil
}

func doDefer() error {
	nums := []int{1}

	// panickingHandler(nums)
	_, err := withRecovery(panickingHandler, nums)
	if err != nil {
		return err
	}
	return nil
}
