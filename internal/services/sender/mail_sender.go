package sender

import (
	"fmt"
	"github.com/wmrodrigues/congress-mailer/internal/structs"
	"time"
)

func SendMails(recipients []structs.Recipient) error {
	for i, item := range recipients {
		fmt.Printf("sending mail %d to %s %s - %s\n", i, item.Treatment, item.Name, item.EmailAddress)
		time.Sleep(time.Millisecond * 100)
	}

	return nil
}
