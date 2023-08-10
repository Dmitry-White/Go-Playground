package smtp

import (
	"bytes"
	"log"

	mail "github.com/xhit/go-simple-mail/v2"
)

func (m *Mailer) BuildHTML(msg Message) (string, error) {
	tmpl, err := render("main.html.gohtml")
	if err != nil {
		log.Println("[BuildHTML] render error", err)
		return "", err
	}

	tmplBuffer := bytes.Buffer{}
	mailErr := tmpl.ExecuteTemplate(&tmplBuffer, "email", msg.DataMap)
	if mailErr != nil {
		return "", mailErr
	}

	rawHTML := tmplBuffer.String()
	renderedHTML, err := inlineCSS(rawHTML)
	if err != nil {
		return "", err
	}

	return renderedHTML, nil
}

func (m *Mailer) BuildPlain(msg Message) (string, error) {
	tmpl, err := render("main.plain.gohtml")
	if err != nil {
		return "", err
	}

	tmplBuffer := bytes.Buffer{}
	mailErr := tmpl.ExecuteTemplate(&tmplBuffer, "email", msg.DataMap)
	if mailErr != nil {
		return "", mailErr
	}

	rawHTML := tmplBuffer.String()

	return rawHTML, nil
}

func (m *Mailer) SendSMTP(msg Message) error {
	if msg.From == "" {
		msg.From = m.FromAddress
	}

	if msg.FromName == "" {
		msg.FromName = m.FromName
	}

	msg.DataMap = map[string]any{
		"message": msg.Data,
	}

	htmlMessage, err := m.BuildHTML(msg)
	if err != nil {
		log.Println("[SendSMTP] BuildHTML error", err)
		return err
	}

	plainMessage, err := m.BuildPlain(msg)
	if err != nil {
		log.Println("[SendSMTP] BuildPlain error", err)
		return err
	}

	log.Println("[SendSMTP] HTML Message: ", htmlMessage)
	log.Println("[SendSMTP] Plain Message: ", plainMessage)

	email := mail.NewMSG()
	email.SetFrom(msg.From).
		AddTo(msg.To).
		SetSubject(msg.Subject).
		SetBody(mail.TextPlain, plainMessage).
		AddAlternative(mail.TextHTML, htmlMessage)

	if len(msg.Attachments) > 0 {
		for _, a := range msg.Attachments {
			email.AddAttachment(a)
		}
	}

	sendErr := email.Send(client)
	if sendErr != nil {
		log.Println("[SendSMTP] Client error", err)
		return sendErr
	}

	return nil
}
