package repo_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_GetAllProducts(t *testing.T) {
	t.Run("get products", func(t *testing.T) {
		rp := initRepo(t)
		products := rp.GetAllProducts()
		assert.Greater(t, len(products), 0)
		assert.Equal(t, existingProduct, products[0].ID)
	})
}
