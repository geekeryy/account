package data

import (
	"gorm.io/gorm"
)

type AccountModel struct {
	gorm.Model
	Name string
}

type AccountRepo interface {
	Get(id int) (*AccountModel, error)
}

func NewAccountRepo(data *Data) AccountRepo {
	return &accountRepo{db: data.Account}
}

type accountRepo struct {
	db *gorm.DB
}

func (w accountRepo) Get(id int) (*AccountModel, error) {
	return &AccountModel{
		Name: "https://www.a.com",
	}, nil
}
