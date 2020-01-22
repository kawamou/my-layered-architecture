package database

import (
	"go-cleanarchitecture-restapi/domain"
)

type MemoRepository struct {
	SqlHandler
}

func (repo *MemoRepository) FindById(id int) (memo domain.Memo, err error) {
	if err = repo.Find(&memo, id).Error; err != nil {
		return
	}
	return
}

func (repo *MemoRepository) FindAll() (memos domain.Memos, err error) {
	if err = repo.Find(&memos).Error; err != nil {
		return
	}
	return
}

func (repo *MemoRepository) Store(m domain.Memo) (memo domain.Memo, err error) {
	if err = repo.Create(&m).Error; err != nil {
		return
	}
	memo = m
	return
}