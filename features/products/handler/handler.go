package handler

import (
	"github.com/dimasyudhana/xyz-multifinance/features/products"
	"github.com/dimasyudhana/xyz-multifinance/utils/response"
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
		res, err := ph.srv.Get(c.QueryParam("keyword"))
		if err != nil {
			jsonResponse, httpCode := response.WebResponseError(err, response.FEAT_PRODUCT_CODE)
			return c.JSON(httpCode, jsonResponse)
		}

		var result []*products.ProductsCore
		for product := range res {
			result = append(result, &products.ProductsCore{
				ProductId: product.ProductId,
				Category:  product.Category,
				SKU:       product.SKU,
				Name:      product.Name,
				Amount:    product.Amount,
				Stock:     product.Stock,
				Tenor:     product.Tenor,
				Bunga:     product.Bunga,
				Fee:       product.Fee,
			})
		}

		mapResponse, httpCode := response.WebResponseSuccess("[success] customer transactions data", response.FEAT_PRODUCT_CODE, result)
		return c.JSON(httpCode, mapResponse)
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
