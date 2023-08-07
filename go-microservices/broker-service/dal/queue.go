package dal

import (
	"log"
	"os"
	"time"

	amqp "github.com/rabbitmq/amqp091-go"
)

var queue *amqp.Connection

var retryCounts int64

func New(queuePool *amqp.Connection) Models {
	queue = queuePool

	return Models{
		Emmiter: Emmiter{},
	}
}

func openAMQP(dsn string) (*amqp.Connection, error) {
	connection, err := amqp.Dial(dsn)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	return connection, nil
}

func Connect() *amqp.Connection {
	dsn := os.Getenv("DSN")

	for {
		connection, err := openAMQP(dsn)
		if err != nil {
			log.Println("AMQP Server is not yet ready...")
			retryCounts++
		} else {
			log.Println("Connected to AMQP Server!")
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
