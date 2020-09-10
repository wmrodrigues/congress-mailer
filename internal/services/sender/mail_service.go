package sender

import (
	"bytes"
	"fmt"
	"github.com/wmrodrigues/congress-mailer/internal/structs"
	"log"
	"net/smtp"
	"os"
	"text/template"
	"time"
)

type Sender struct {
	settings structs.Settings
	recipients []structs.Recipient
}

// NewSender creates a new instance of Sender struct
func NewSender(settingsFile structs.Settings) *Sender {
	return &Sender{settings: settingsFile}
}

// SetRecipients set the recipients attribute
func (s *Sender) SetRecipients(r []structs.Recipient) {
	s.recipients = r
}

func getMessageTemplate() *template.Template {
	wd, err := os.Getwd()
	if err != nil {
		err = fmt.Errorf("error getting working dir for template file, %s", err.Error())
		log.Fatal(err)
	}

	templateFilePath := fmt.Sprintf("%s/configs/message.template", wd)
	_template := template.Must(template.ParseFiles(templateFilePath))
	return _template
}

func (s *Sender) sendMail(recipient structs.Recipient) {
	_smtp := s.settings.Smtp
	_mail := s.settings.Mail

	_template := getMessageTemplate()
	var content bytes.Buffer
	err := _template.Execute(&content, recipient)

	if err != nil {
		err = fmt.Errorf("error loading content on template, %s", err.Error())
		log.Fatal(err)
	}

	date := time.Now().Format("2006-01-02")
	subject := fmt.Sprintf(_mail.Subject, date)

	message := []byte(fmt.Sprintf("To: %s\r\n", recipient.EmailAddress) +
		fmt.Sprintf("Subject: %s\r\n", subject) +
		fmt.Sprintf("From: %s\r\n", _smtp.From) +
		"\r\n" +
		content.String())

	to := []string{recipient.EmailAddress}
	auth := smtp.PlainAuth("", _smtp.Username, _smtp.Password, _smtp.Host)
	smtAddress := fmt.Sprintf("%s:%d", _smtp.Host, _smtp.Port)
	err = smtp.SendMail(smtAddress, auth, _smtp.From, to, message)

	if err != nil {
		err = fmt.Errorf("error sending mail to %s, %s", recipient.EmailAddress, err.Error())
		log.Println(err)
		return
	}

	log.Printf("mail succesfully sent to %s", recipient.EmailAddress)
}

// SendMails sends the mail message to all mail addresses in the Recipient attribute
func (s *Sender)SendMails() error {
	for i, item := range s.recipients {
		fmt.Printf("sending mail %d to %s %s - %s\n", i, item.Treatment, item.Name, item.EmailAddress)
		s.sendMail(item)
		time.Sleep(time.Millisecond * 100)
	}

	return nil
}
