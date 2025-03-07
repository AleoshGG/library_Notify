package infrastructure

import "library-Notify/src/infrastructure/adapters"

var notifier *adapters.Notifier
var fetchAPI *adapters.FetchAPI

func GoDependences() {
	notifier = adapters.NewNotifier()
	fetchAPI = adapters.NewFetchAPI()
}

func GetNotifier() *adapters.Notifier {
	return notifier
}

func GetFeychAPI() *adapters.FetchAPI {
	return fetchAPI
}