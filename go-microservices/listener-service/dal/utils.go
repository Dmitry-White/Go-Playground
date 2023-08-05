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

func declareQueue(ch *amqp.Channel, name string) (amqp.Queue, error) {
	return ch.QueueDeclare(
		name,  // name?
		false, // durable?
		false, // delete when unused?
		true,  // exclusive?
		false, // no-wait?
		nil,   // arguments?
	)
}
