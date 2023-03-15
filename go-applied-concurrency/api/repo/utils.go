package repo

import (
	"fmt"
	"go-applied-concurrency/api/models"
	"math"
)

// validateItem runs validations on a given order
func (r *Repo) validateItem(item models.Item) error {
	if item.Amount < 1 {
		return fmt.Errorf("order amount must be at least 1:got %d", item.Amount)
	}
	if err := r.Products.Exists(item.ProductID); err != nil {
		return fmt.Errorf("product %s does not exist", item.ProductID)
	}
	return nil
}

// processOrder is an internal method which completes or rejects an order
func processOrder(r *Repo, order *models.Order) {
	item := order.Item
	product, err := r.Products.Find(item.ProductID)
	if err != nil {
		order.Status = models.OrderStatus_Rejected
		order.Error = err.Error()
		return
	}
	if product.Stock < item.Amount {
		order.Status = models.OrderStatus_Rejected
		order.Error = fmt.Sprintf("not enough stock for product %s:got %d, want %d", item.ProductID, product.Stock, item.Amount)
		return
	}
	remainingStock := product.Stock - item.Amount
	product.Stock = remainingStock
	r.Products.Upsert(product)

	total := math.Round(float64(order.Item.Amount)*product.Price*100) / 100
	order.Total = &total
	order.Complete()
}

// Mutex essentially adds state to an API
// which is inherently flawed.
// func ProcessOrders(r *Repo, order *models.Order) {
// 	r.Mutex.Lock()
// 	processOrder(r, order)
// 	r.Orders.Upsert(*order)
// 	r.Mutex.Unlock()
// 	fmt.Printf("Processing order %s completed\n", order.ID)
// }

// Simply closing Incoming channel will not be enough
// to ensure we're not reading race conditionally
// from an already closed channel
// func ProcessOrders(r *Repo /*order *models.Order*/) {
// 	fmt.Println("Order processing started!")
// 	for order := range r.Incoming {
// 		processOrder(r, &order)
// 		r.Orders.Upsert(order)
// 		fmt.Printf("Processing order %s completed\n", order.ID)
// 	}
// 	fmt.Println("Order processing stopped!")
// }

func ProcessOrders(r *Repo /*order *models.Order*/) {
	fmt.Println("Order processing started!")
	for {
		select {
		case order := <-r.Incoming:
			processOrder(r, &order)
			r.Orders.Upsert(order)
			r.Processed <- order
			fmt.Printf("Processing order %s completed\n", order.ID)
		case <-r.Done:
			fmt.Println("Order processing stopped!")
			return
		}
	}
}
