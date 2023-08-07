package dal

import (
	"context"

	amqp "github.com/rabbitmq/amqp091-go"
)

func (e *Emmiter) SetupExchange(name string) error {
	channel, err := queue.Channel()
	if err != nil {
		return err
	}
	defer channel.Close()

	exchangeErr := declareExchange(channel, name)
	if exchangeErr != nil {
		return exchangeErr
	}

	e.Exchange = name

	return nil
}

func (e *Emmiter) Push(key string, payload []byte) error {
	channel, err := queue.Channel()
	if err != nil {
		return err
	}
	defer channel.Close()

	ctx := context.Background()
	msg := amqp.Publishing{
		ContentType: "text/plain",
		Body:        payload,
	}

	publishErr := channel.PublishWithContext(
		ctx,
		e.Exchange,
		key,
		false,
		false,
		msg,
	)
	if publishErr != nil {
		return publishErr
	}

	return nil
}
