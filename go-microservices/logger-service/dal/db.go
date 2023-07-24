package dal

import (
	"context"
	"log"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var db *mongo.Client

var retryCounts int64

// New is the function used to create an instance of the dal package. It returns the type
// Model, which embeds all the types we want to be available to our application.
func New(dbPool *mongo.Client) interface{} {
	db = dbPool

	return db
}

func openDB(uri, username, password string) (*mongo.Client, error) {
	clientOptions := options.Client().ApplyURI(uri)
	optionsCredentials := options.Credential{
		Username: username,
		Password: password,
	}

	clientOptions.SetAuth(optionsCredentials)

	db, err := mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		return nil, err
	}

	return db, nil
}

func Connect() *mongo.Client {
	uri := os.Getenv("URI")
	username := os.Getenv("Username")
	password := os.Getenv("Password")

	for {
		connection, err := openDB(uri, username, password)
		if err != nil {
			log.Println("DB is not yet ready...")
			retryCounts++
		} else {
			log.Println("Connected to DB!")
			return connection
		}

		if retryCounts > RETRY_LIMIT {
			log.Println(err)
			return nil
		}

		log.Printf("Backing off for %d seconds...\n", RETRY_TIMEOUT/time.Second)
		time.Sleep(RETRY_TIMEOUT)
		continue
	}
}
