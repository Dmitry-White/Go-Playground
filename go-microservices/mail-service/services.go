package main

import (
	"go-microservices/mail-service/smtp"
	"log"
)

func (s *Services) send(requestPayload *RequestPayload) (*smtp.Message, error) {
	msg := smtp.Message{
		From:    requestPayload.From,
		To:      requestPayload.To,
		Subject: requestPayload.Subject,
		Data:    requestPayload.Message,
	}

	err := s.Models.Mailer.SendSMTP(msg)
	if err != nil {
		log.Println("[send] Model Error: ", err)
		return nil, err
	}

	return &msg, nil
}
