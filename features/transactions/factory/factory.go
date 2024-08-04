package factory

import (
	"database/sql"

	_transactionData "github.com/dimasyudhana/xyz-multifinance/features/transactions/data"
	_transactionHandler "github.com/dimasyudhana/xyz-multifinance/features/transactions/handler"
	_transactionService "github.com/dimasyudhana/xyz-multifinance/features/transactions/service"
)

func New(db *sql.DB) *_transactionHandler.TransactionHandler {
	data := _transactionData.New(db)
	service := _transactionService.New(data)
	handler := _transactionHandler.New(service)
	return handler
}
