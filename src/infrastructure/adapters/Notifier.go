package adapters

import "fmt"

type Notifier struct{}

func NewNotifier() *Notifier {
	return &Notifier{}
}

func (n *Notifier) NotifyOfLend(msg string, email string, return_date string, name string) {
	fmt.Print(msg+email+return_date+name)
}

func (n *Notifier) NotifyOfReturn(msg string, email string, name string) {
	fmt.Print(msg+email+name)
}