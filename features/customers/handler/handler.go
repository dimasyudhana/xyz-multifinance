package handler

import (
	"errors"

	"github.com/dimasyudhana/xyz-multifinance/features/customers"
	"github.com/dimasyudhana/xyz-multifinance/utils/encrypt"
	"github.com/dimasyudhana/xyz-multifinance/utils/response"
	"github.com/labstack/echo/v4"
)

type CustomerHandler struct {
	srv customers.Service
}

func New(srv customers.Service) *CustomerHandler {
	return &CustomerHandler{
		srv: srv,
	}
}

func (ch *CustomerHandler) CustomerTransactions() echo.HandlerFunc {
	return func(c echo.Context) error {
		customerId, err := encrypt.ExtractToken(c)
		if err != nil {
			jsonResponse, httpCode := response.WebResponseError(errors.New(response.JWT_FailedCastingJwtToken), response.FEAT_CUSTOMER_CODE)
			return c.JSON(httpCode, jsonResponse)
		}

		res, err := ch.srv.CustomerTransactions(customerId)
		if err != nil {
			jsonResponse, httpCode := response.WebResponseError(err, response.FEAT_CUSTOMER_CODE)
			return c.JSON(httpCode, jsonResponse)
		}

		var result []*customers.CustomersCore
		for customer := range res {
			result = append(result, customer)
		}

		trx := []*map[string]any{}
		for _, customer := range result {
			trx = append(trx, &map[string]any{
				"nomor_kontrak":    customer.NomorKontrak,
				"nama_asset":       customer.NamaAsset,
				"otr":              customer.OTR,
				"admin_fee":        customer.AdminFee,
				"jumlah_cicilan":   customer.JumlahCicilan,
				"jumlah_bunga":     customer.JumlahBunga,
				"transaction_date": customer.TransactionDate,
			})
		}

		customerTransactions := map[string]any{
			"customer_id":  result[0].CustomerId,
			"email":        result[0].Email,
			"phone":        result[0].Phone,
			"fullname":     result[0].Fullname,
			"transactions": trx,
		}

		mapResponse, httpCode := response.WebResponseSuccess("[success] customer transactions data", response.FEAT_CUSTOMER_CODE, customerTransactions)
		return c.JSON(httpCode, mapResponse)
	}
}
