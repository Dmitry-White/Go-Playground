package smtp

type Models struct {
	Mailer Mailer
}

type Mailer struct {
	Domain      string
	Host        string
	Port        int
	Username    string
	Password    string
	Encryption  string
	FromAddress string
	FromName    string
}

type Message struct {
	From        string
	FromName    string
	To          string
	Subject     string
	Attachments []string
	Data        any
	DataMap     map[string]any
}

type Paths struct {
	EMAILS    string
	TEMPLATES string
}
