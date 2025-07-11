package repository

import (
	"github.com/briandang59/be_scada/config"
	"github.com/briandang59/be_scada/internal/model"
)

type AccountRepository interface {
	FindByUsername(username string) (*model.Account, error)
	Create(account *model.Account) error
}

type accountRepo struct{}

func NewAccountRepository() AccountRepository {
	return &accountRepo{}
}

func (r *accountRepo) FindByUsername(username string) (*model.Account, error) {
	var account model.Account
	if err := config.DB.Where("username = ?", username).First(&account).Error; err != nil {
		return nil, err
	}
	return &account, nil
}

func (r *accountRepo) Create(account *model.Account) error {
	return config.DB.Create(account).Error
}
