package dal

import amqp "github.com/rabbitmq/amqp091-go"

func (c *Consumer) SetupExchange(name string) error {
	channel, err := queue.Channel()
	if err != nil {
		return err
	}

	exchangeErr := declareExchange(channel, name)
	if exchangeErr != nil {
		return exchangeErr
	}

	c.Exchange = name

	return nil
}

func (c *Consumer) SetupQueue(name string) error {
	channel, err := queue.Channel()
	if err != nil {
		return err
	}

	queue, err := declareQueue(channel, name)
	if err != nil {
		return err
	}

	c.Queue = &queue

	return nil
}

func (c *Consumer) Subscribe(topics []string) error {
	channel, err := queue.Channel()
	if err != nil {
		return err
	}

	for _, topic := range topics {
		channel.QueueBind(
			c.Queue.Name,
			topic,
			c.Exchange,
			false,
			nil,
		)

		if err != nil {
			return err
		}
	}

	return nil
}

func (c *Consumer) Listen() (<-chan amqp.Delivery, error) {
	channel, err := queue.Channel()
	if err != nil {
		return nil, err
	}

	messages, err := channel.Consume(c.Queue.Name, "", true, false, false, false, nil)
	if err != nil {
		return nil, err
	}

	return messages, nil
}
