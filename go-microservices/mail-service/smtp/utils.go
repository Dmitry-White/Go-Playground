package smtp

import (
	"fmt"
	"html/template"
	"os"

	"github.com/vanng822/go-premailer/premailer"
	mail "github.com/xhit/go-simple-mail/v2"
)

func render(t string) (*template.Template, error) {
	files, err := os.ReadDir("./emails")
	if err != nil {
		return nil, err
	}
	for _, f := range files {
		fmt.Println("File:", f.Name())
	}

	templateSlice := []string{}
	templateSlice = append(templateSlice, fmt.Sprintf("%s/%s", PATHS.EMAILS, t))

	partials := []string{
		PATHS.TEMPLATES + "/base.layout.gohtml",
		PATHS.TEMPLATES + "/header.partial.gohtml",
		PATHS.TEMPLATES + "/footer.partial.gohtml",
	}
	templateSlice = append(templateSlice, partials...)

	tmpl, err := template.ParseFiles(templateSlice...)
	if err != nil {
		return nil, err
	}

	return tmpl, nil
}

func inlineCSS(document string) (string, error) {
	options := premailer.Options{
		RemoveClasses:     false,
		CssToAttributes:   false,
		KeepBangImportant: true,
	}

	prem, err := premailer.NewPremailerFromString(document, &options)
	if err != nil {
		return "", err
	}

	html, err := prem.Transform()
	if err != nil {
		return "", err
	}

	return html, nil
}

func getEncryption(s string) mail.Encryption {
	switch s {
	case "tls":
		return mail.EncryptionSTARTTLS
	case "ssl":
		return mail.EncryptionSSLTLS
	case "none", "":
		return mail.EncryptionNone
	default:
		return mail.EncryptionSTARTTLS
	}
}
