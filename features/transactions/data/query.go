package data

import (
	"database/sql"

	"github.com/dimasyudhana/xyz-multifinance/features/transactions"
)

type transactionQuery struct {
	db *sql.DB
}

func New(db *sql.DB) transactions.Data {
	return &transactionQuery{
		db: db,
	}
}

func (t *transactionQuery) Create(data *transactions.TransactionProductCore) (*transactions.TransactionCore, error) {
	panic("unimplemented")
}

func (t *transactionQuery) CheckStatus(data *transactions.TransactionProductCore) (*transactions.TransactionCore, error) {
	panic("unimplemented")
}

func (t *transactionQuery) GetById(transactionId string) (*transactions.TransactionCore, error) {
	panic("unimplemented")
}
