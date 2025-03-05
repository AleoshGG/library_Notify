package services

import "library-Notify/src/application/repository"

type NotifyOfReturnEvent struct {
	n repository.INotify
}

func NewNotifyOfReturnEvent(n repository.INotify) *NotifyOfReturnEvent{
	return &NotifyOfReturnEvent{n: n}
}

func (s *NotifyOfReturnEvent) Run(msg string, email string, name string) {
	s.n.NotifyOfReturn(msg, email, name)
}