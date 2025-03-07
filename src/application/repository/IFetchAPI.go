package repository

import "library-Notify/src/domain"

type IFetchAPI interface {
	FetchAPI(id_reader int) (domain.Response, error)
}