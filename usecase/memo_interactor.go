package usecase

import (
	"go-cleanarchitecture-restapi/domain"
)

type MemoInteractor struct {
	MemoRepository MemoRepository
}

func (interactor *MemoInteractor) MemoById(id int) (memo domain.Memo, err error) {
	memo, err = interactor.MemoRepository.FindById(id)
	return
}

func (interactor *MemoInteractor) Memos() (memos domain.Memos, err error) {
	memos, err = interactor.MemoRepository.FindAll()
	return
}

func (interactor *MemoInteractor) Add(m domain.Memo) (memo domain.Memo, err error) {
	memo, err = interactor.MemoRepository.Store(m)
	return
}