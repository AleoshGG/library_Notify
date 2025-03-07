package services

import (
	"library-Notify/src/application/repository"
	"library-Notify/src/domain"
)

type FetchAPIService struct {
	f repository.IFetchAPI
}

func NewFetchAPIService(f repository.IFetchAPI) *FetchAPIService {
	return &FetchAPIService{f: f}
}

func (s *FetchAPIService) Run(id_reader int) (domain.Response, error) {
	return s.f.FetchAPI(id_reader) 
}