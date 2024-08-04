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

func (p *productService) Get(data *products.ProductsCode) (*[]products.ProductsCode, error) {
	panic("unimplemented")
}

func (p *productService) GetById(productId string) (*products.ProductsCode, error) {
	panic("unimplemented")
}

func (p *productService) CheckAvailabitity(data *products.ProductsCode) (*products.ProductsCode, error) {
	panic("unimplemented")
}
