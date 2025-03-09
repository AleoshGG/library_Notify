package adapters

import (
	"fmt"
	"os"

	"github.com/resend/resend-go/v2"
)

type Notifier struct{}

func NewNotifier() *Notifier {
	return &Notifier{}
}

func (n *Notifier) NotifyOfLend(msg string, email string, return_date string, name string) {
    client := resend.NewClient(os.Getenv("APIKEY"))

    params := &resend.SendEmailRequest{
        From:    "onboarding@resend.dev",
        To:      []string{email},
        Html:    "<strong>Hola "+name+"</strong>"+"<p>"+msg+return_date+"</p>",
        Subject: "Estado de préstamo",
    }

    sent, err := client.Emails.Send(params)
    if err != nil {
        fmt.Println(err.Error())
        return
    }
    fmt.Println(sent.Id)
}

func (n *Notifier) NotifyOfReturn(msg string, email string, name string) {
	client := resend.NewClient(os.Getenv("APIKEY"))

    params := &resend.SendEmailRequest{
        From:    "onboarding@resend.dev",
        To:      []string{email},
        Html:    "<strong>Hola "+name+"<strong>"+"<p>"+msg+"</p>",
        Subject: "Estado de la devolución",
    }

    sent, err := client.Emails.Send(params)
    if err != nil {
        fmt.Println(err.Error())
        return
    }
    fmt.Println(sent.Id)
}