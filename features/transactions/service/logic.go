package service

import (
	"github.com/dimasyudhana/xyz-multifinance/features/transactions"
	"github.com/go-playground/validator/v10"
)

type transactionService struct {
	transactionData transactions.Data
	validate        *validator.Validate
}

func New(transactionData transactions.Data) transactions.Service {
	return &transactionService{
		transactionData: transactionData,
		validate:        validator.New(),
	}
}

func (t *transactionService) Create(data *transactions.TransactionProductCore) (*transactions.TransactionCore, error) {
	panic("unimplemented")
}

func (t *transactionService) CheckStatus(data *transactions.TransactionProductCore) (*transactions.TransactionCore, error) {
	panic("unimplemented")
}

func (t *transactionService) GetById(transactionId string) (*transactions.TransactionCore, error) {
	panic("unimplemented")
}
