package repo_test

import (
	"fmt"
	"go-applied-concurrency/api/db"
	"go-applied-concurrency/api/models"
	"go-applied-concurrency/api/repo"
	"sync"
	"testing"

	"github.com/stretchr/testify/assert"
)

const productCode = "TEST"
const productStock = 11

// how many goroutines we will place orders on
const concurrentOrders = 10

func assertStock(t *testing.T, r *repo.Repo, expectedStock int) {
	prod, err := r.Products.Find(productCode)
	assert.Nil(t, err)
	assert.Equal(t, expectedStock, prod.Stock)
}

// THIS TEST IS FLAKY. FOR DEMO PURPOSES ONLY
func Test_ProcessOrder(t *testing.T) {
	// Uncomment out line below to skip it
	// t.Skip("Skipping process Order test")

	prod := &db.ProductDB{}
	prod.Upsert(models.Product{
		ID:    productCode,
		Stock: productStock,
	})
	r := &repo.Repo{
		Products: prod,
		Orders:   db.NewOrders(),
	}
	item := models.Item{
		ProductID: productCode,
		Amount:    1,
	}

	go repo.ProcessOrders(r)

	t.Run(fmt.Sprintf("%d concurrent orders", concurrentOrders), func(t *testing.T) {
		var wg sync.WaitGroup
		wg.Add(concurrentOrders)
		for j := 0; j < concurrentOrders; j++ {
			go func(wg *sync.WaitGroup) {
				defer wg.Done()
				order := models.NewOrder(item)
				// repo.ProcessOrders(r, &order)
				r.Incoming <- order
			}(&wg)
		}
		wg.Wait()
		expected := productStock - concurrentOrders
		assertStock(t, r, expected)
	})
}
