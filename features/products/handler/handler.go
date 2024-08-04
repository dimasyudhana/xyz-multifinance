package handler

import (
	"github.com/dimasyudhana/xyz-multifinance/features/products"
	"github.com/labstack/echo/v4"
)

type ProductHandler struct {
	srv products.Service
}

func New(srv products.Service) *ProductHandler {
	return &ProductHandler{
		srv: srv,
	}
}

func (ph *ProductHandler) Get() echo.HandlerFunc {
	return func(c echo.Context) error {
		panic("belum di implementasi")
	}
}

func (ph *ProductHandler) GetById() echo.HandlerFunc {
	return func(c echo.Context) error {
		panic("belum di implementasi")
	}
}

func (ph *ProductHandler) CheckAvailabitity() echo.HandlerFunc {
	return func(c echo.Context) error {
		panic("belum di implementasi")
	}
}
