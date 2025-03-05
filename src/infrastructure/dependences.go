package infrastructure

var notifier *Notifier

func GoDependences() {
	notifier = NewNotifier()
}

func GetNotifier() *Notifier {
	return notifier
}