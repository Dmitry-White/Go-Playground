package main

import (
	"encoding/json"
	"log"

	amqp "github.com/rabbitmq/amqp091-go"
)

func (app *AppConfig) listen() error {
	err := app.Services.listen(app.handleMessages)
	if err != nil {
		return err
	}

	return nil
}

func (app *AppConfig) handleTopics() error {
	exchangeErr := app.Services.Models.Consumer.SetupExchange(EXCHANGES.LOGS)
	if exchangeErr != nil {
		return exchangeErr
	}

	queueErr := app.Services.Models.Consumer.SetupQueue(QUEUES.RANDOM)
	if queueErr != nil {
		return queueErr
	}

	topics := []string{
		TOPICS.LOG_INFO,
		TOPICS.LOG_WARNING,
		TOPICS.LOG_ERROR,
	}
	err := app.Services.Models.Consumer.Subscribe(topics)
	if err != nil {
		return err
	}

	return nil
}

func (app *AppConfig) handleMessages(messages <-chan amqp.Delivery) {
	for message := range messages {
		log.Println("[handleMessages] Recieved a message, processing...")

		payload := LogPayload{}
		err := json.Unmarshal(message.Body, &payload)
		if err != nil {
			log.Println("[handleMessages] Unmarshal Error:", err)
		}

		go app.handlePayload(payload)
	}
}

func (app *AppConfig) handlePayload(payload LogPayload) {
	switch payload.Name {
	case EVENTS.LOG:
		err := app.Services.log(payload)
		if err != nil {
			log.Println("[handlePayload] Log Error:", err)
		}
	default:
		err := app.Services.log(payload)
		if err != nil {
			log.Println("[handlePayload] Default Error:", err)
		}
	}
}
