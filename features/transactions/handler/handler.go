package handler

import (
	"github.com/dimasyudhana/xyz-multifinance/features/transactions"
	"github.com/labstack/echo/v4"
)

type TransactionHandler struct {
	srv transactions.Service
}

func New(srv transactions.Service) *TransactionHandler {
	return &TransactionHandler{
		srv: srv,
	}
}

func (ph *TransactionHandler) Create() echo.HandlerFunc {
	return func(c echo.Context) error {
		panic("belum di implementasi")
	}
}

func (ph *TransactionHandler) CheckStatus() echo.HandlerFunc {
	return func(c echo.Context) error {
		panic("belum di implementasi")
	}
}

func (ph *TransactionHandler) GetById() echo.HandlerFunc {
	return func(c echo.Context) error {
		panic("belum di implementasi")
	}
}
