package smtp

import (
	"log"
	"os"
	"strconv"
	"time"

	mail "github.com/xhit/go-simple-mail/v2"
)

var client *mail.SMTPClient

var retryCounts int64

// New is the function used to create an instance of the smtp package. It returns the type
// Models, which embeds all the types we want to be available to our application.
func New(clientPool *mail.SMTPClient, mailer *Mailer) Models {
	client = clientPool

	return Models{
		Mailer: *mailer,
	}
}

func openSMTP(m *Mailer) (*mail.SMTPClient, error) {
	server := mail.SMTPServer{
		Host:           m.Host,
		Port:           m.Port,
		Username:       m.Username,
		Password:       m.Password,
		Encryption:     getEncryption(m.Encryption),
		KeepAlive:      false,
		ConnectTimeout: 10 * time.Second,
		SendTimeout:    10 * time.Second,
	}

	connection, err := server.Connect()
	if err != nil {
		log.Println(err)
		return nil, err
	}

	return connection, nil
}

func Connect() (*mail.SMTPClient, *Mailer) {
	host := os.Getenv("Host")
	port, _ := strconv.Atoi(os.Getenv("Port"))
	username := os.Getenv("Username")
	password := os.Getenv("Password")
	encryption := os.Getenv("Encryption")

	mailer := Mailer{
		Host:       host,
		Port:       port,
		Username:   username,
		Password:   password,
		Encryption: encryption,
	}

	for {
		connection, err := openSMTP(&mailer)
		if err != nil {
			log.Println("SMTP Server is not yet ready...")
			retryCounts++
		} else {
			log.Println("Connected to SMTP Server!")
			return connection, &mailer
		}

		if retryCounts > RETRY_LIMIT {
			log.Println(err)
			return nil, nil
		}

		log.Printf("Backing off for %d seconds...\n", RETRY_TIMEOUT/time.Second)
		time.Sleep(RETRY_TIMEOUT)
		continue
	}
}
