package infrastructure

type Notifier struct{}

func NewNotifier() *Notifier {
	return &Notifier{}
}

func (n *Notifier) NotifyOfLend(msg string, email string, return_date string, name string) {

}

func (n *Notifier) NotifyOfReturn(msg string, email string, name string) {

}