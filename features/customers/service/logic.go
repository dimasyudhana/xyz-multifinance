package service

import (
	"github.com/dimasyudhana/xyz-multifinance/features/customers"
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

func (c *customerService) CustomerTransactions(customerId string) (<-chan *customers.CustomersCore, error) {
	rows, err := c.customerData.CustomerTransactions(customerId)
	if err != nil {
		return nil, err
	}

	result := make(chan *customers.CustomersCore)
	go func() {
		for _, data := range rows {
			result <- data
		}
		close(result)
	}()

	return result, nil
}
