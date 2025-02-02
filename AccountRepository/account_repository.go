package AccountRepository

import (
	"golangmicro/models"
	"gorm.io/gorm"
)

type TransferRepository interface {
	CreateBalance(balance int, accountId int) (*models.Account, error)
	TransferFunds(balance int, fromAccountId int, toAccountId int) error
}

type transferRepository struct {
	DB *gorm.DB
}

func NewTransferRepository(db *gorm.DB) TransferRepository {
	return &transferRepository{DB: db}
}

func (t *transferRepository) CreateBalance(balance int, accountId int) (*models.Account, error) {

}

func (t *transferRepository) TransferFunds(balance int, fromAccountId int, toAccountId int) error {
}
