package main

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
