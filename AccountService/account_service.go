package AccountService

import (
	"golangmicro/AccountRepository"
	"golangmicro/models"
)

type ServiceAccount interface {
	CreateBalance(balance int, accountId int) (*models.Account, error)
	TransferFunds(balance int, fromAccountId int, toAccountId int) error
}

type serviceAccount struct {
	accountRepo AccountRepository.TransferRepository
}

func NewServiceRepository(accountRepo AccountRepository.TransferRepository) ServiceAccount {
	return &serviceAccount{accountRepo: accountRepo}
}

func (nsr *serviceAccount) CreateBalance(balance int, accountId int) (*models.Account, error) {

}

func (nsr *serviceAccount) TransferFunds(balance int, fromAccountId int, toAccountId int) error {

}
