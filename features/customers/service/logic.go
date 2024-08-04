package service

import (
	"errors"

	"github.com/dimasyudhana/xyz-multifinance/features/customers"
	"github.com/dimasyudhana/xyz-multifinance/utils/encrypt"
	"github.com/go-playground/validator/v10"
)

type customerService struct {
	customerData customers.Data
	validate     *validator.Validate
}

func New(customerData customers.Data) customers.Service {
	return &customerService{
		customerData: customerData,
		validate:     validator.New(),
	}
}

func (cs *customerService) Insert(data *customers.CustomersCore) (*customers.CustomersCore, error) {
	hash, err := encrypt.HashPassword(data.Password)
	if err != nil {
		return nil, errors.New(err.Error())
	}
	data.Password = hash
	result, err := cs.customerData.Insert(data)
	if err != nil {
		return nil, errors.New(err.Error())
	}
	return result, nil
}
