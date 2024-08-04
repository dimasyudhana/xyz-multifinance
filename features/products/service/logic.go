package service

import (
	"github.com/dimasyudhana/xyz-multifinance/features/products"
	"github.com/go-playground/validator/v10"
)

type productService struct {
	productData products.Data
	validate    *validator.Validate
}

func New(productData products.Data) products.Service {
	return &productService{
		productData: productData,
		validate:    validator.New(),
	}
}

func (p *productService) Get(keyword string) (<-chan *products.ProductsCore, error) {

	rows, err := p.productData.Get(keyword)
	if err != nil {
		return nil, err
	}

	result := make(chan *products.ProductsCore)
	go func() {
		for _, row := range rows {
			result <- row
		}
		close(result)
	}()

	return result, nil
}

func (p *productService) GetById(productId string) (*products.ProductsCore, error) {
	panic("unimplemented")
}

func (p *productService) CheckAvailabitity(data *products.ProductsCore) (*products.ProductsCore, error) {
	panic("unimplemented")
}
