package factory

import (
	"database/sql"

	_authData "github.com/dimasyudhana/xyz-multifinance/features/authentication/data"
	_authHandler "github.com/dimasyudhana/xyz-multifinance/features/authentication/handler"
	_authService "github.com/dimasyudhana/xyz-multifinance/features/authentication/service"
)

func New(db *sql.DB) *_authHandler.AuthHandler {
	data := _authData.New(db)
	service := _authService.New(data)
	handler := _authHandler.New(service)
	return handler
}
