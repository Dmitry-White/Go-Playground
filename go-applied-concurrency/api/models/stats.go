package models

import "math"

type Statistics struct {
	CompletedOrders int     `json:"completedOrders"`
	RejectedOrders  int     `json:"rejectedOrders"`
	ReversedOrders  int     `json:"reversedOrders"`
	Revenue         float64 `json:"revenue"`
}

// Combine adds the numbers frin a two statistics objects
func Combine(this, that Statistics) Statistics {
	return Statistics{
		CompletedOrders: this.CompletedOrders + that.CompletedOrders,
		ReversedOrders:  that.ReversedOrders + that.ReversedOrders,
		RejectedOrders:  that.RejectedOrders + that.RejectedOrders,
		Revenue:         math.Round((this.Revenue+that.Revenue)*100) / 100,
	}
}
