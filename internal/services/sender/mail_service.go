package sender

import (
	"fmt"
	"github.com/wmrodrigues/congress-mailer/internal/structs"
	"log"
	"net/smtp"
	"time"
)

type Sender struct {
	settings structs.Settings
	recipients []structs.Recipient
}

// NewSender
func NewSender(settingsFile structs.Settings) *Sender {
	return &Sender{settings: settingsFile}
}

// SetRecipients
func (s *Sender) SetRecipients(r []structs.Recipient) {
	s.recipients = r
}

func (s *Sender) sendMail(recipient structs.Recipient) {
	_smtp := s.settings.Smtp

	message := []byte(fmt.Sprintf("To: %s\r\n", recipient.EmailAddress) +
		"Subject: Solicitação importante\r\n" +
		fmt.Sprintf("From: %s\r\n", _smtp.From) +
		"\r\n" +
		fmt.Sprintf("%s %s, bom dia! \nVenho através deste exercer o meu direito de cidadão solicitar que V.Ex.a assine a emenda anti privilegio do Exmo. Senhor Deputado Kim Kataguiri.\r\n\n\n\nMinha sincera gratidão por Vossa atenção", recipient.Treatment, recipient.Name))

	to := []string{recipient.EmailAddress}
	auth := smtp.PlainAuth("", _smtp.Username, _smtp.Password, _smtp.Host)
	smtAddres := fmt.Sprintf("%s:%d", _smtp.Host, _smtp.Port)
	err := smtp.SendMail(smtAddres, auth, _smtp.From, to, message)

	if err != nil {
		fmt.Errorf("error sending mail to %s, %s", recipient.EmailAddress, err.Error())
		log.Println(err)
		return
	}

	log.Printf("mail succesfully sent to %s", recipient.EmailAddress)
}

// SendMails
func (s *Sender)SendMails() error {
	for i, item := range s.recipients {
		fmt.Printf("sending mail %d to %s %s - %s\n", i, item.Treatment, item.Name, item.EmailAddress)
		s.sendMail(item)
		time.Sleep(time.Millisecond * 100)
	}

	return nil
}
