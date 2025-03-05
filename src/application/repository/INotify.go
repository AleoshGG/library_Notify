package repository

type INotify interface {
	NotifyOfLend(msg string, email string, return_date string, name string)
	NotifyOfReturn(msg string, email string, name string)
}