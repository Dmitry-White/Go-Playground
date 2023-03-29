package services

type Message struct {
	From User
	To   User
	Text string
}

func GetMessages(user *User) []Message {
	// Dummy Implementation
	messages := []Message{}

	messages = append(messages, Message{
		From: User{
			Login: "Alice",
		},
		To:   *user,
		Text: "We know",
	})

	return messages
}
