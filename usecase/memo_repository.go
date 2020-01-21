package usecase

import (
	"go-cleanarchitecture-restapi/domain"
)

type MemoRepository interface {
	FindById (id int) (domain.Memo, error)
	FindAll () (domain.Memo, error)
	Store (domain.Memo) (domain.Memo, error)
}