package dal

import amqp "github.com/rabbitmq/amqp091-go"

func declareExchange(ch *amqp.Channel, name string) error {
	return ch.ExchangeDeclare(
		name,    // name
		"topic", // type
		true,    // durable?
		false,   // auto-deleted?
		false,   // internal?
		false,   // no-wait?
		nil,     // arguements?
	)
}
