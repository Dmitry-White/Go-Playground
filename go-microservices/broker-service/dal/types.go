package dal

import amqp "github.com/rabbitmq/amqp091-go"

type Emmiter struct {
	Conn     *amqp.Connection
	Queue    *amqp.Queue
	Exchange string
}

type Models struct {
	Emmiter Emmiter
}
