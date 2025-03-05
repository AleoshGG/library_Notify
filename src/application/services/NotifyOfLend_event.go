package services

import "library-Notify/src/application/repository"

type NotifyOfLendEvent struct {
	n repository.INotify
}

func NewNotifyOfLendEvent(n repository.INotify) *NotifyOfLendEvent {
	return &NotifyOfLendEvent{n: n}
}

func (s *NotifyOfLendEvent) Run(msg string, email string, return_date string, name string) {
	s.n.NotifyOfLend(msg, email, return_date, name)
}