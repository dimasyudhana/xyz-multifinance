package factory

import (
	"database/sql"

	_customerData "github.com/dimasyudhana/xyz-multifinance/features/customers/data"
	_customerHandler "github.com/dimasyudhana/xyz-multifinance/features/customers/handler"
	_customerService "github.com/dimasyudhana/xyz-multifinance/features/customers/service"
)

func New(db *sql.DB) *_customerHandler.CustomerHandler {
	data := _customerData.New(db)
	service := _customerService.New(data)
	handler := _customerHandler.New(service)
	return handler
}
