package repo_test

import (
	"os"
	"testing"

	"go-applied-concurrency/api/db"
	"go-applied-concurrency/api/repo"

	"github.com/stretchr/testify/assert"
)

const existingProduct = "MWBLU"

func initRepo(t *testing.T) repo.IRepo {
	rp, err := repo.New(db.ProductInputPathTest)
	assert.Nil(t, err)
	return rp
}

func TestMain(m *testing.M) {
	if err := os.Chdir(".."); err != nil {
		panic(err)
	}
	code := m.Run()
	os.Exit(code)
}
