package data

import (
	"database/sql"

	"github.com/dimasyudhana/xyz-multifinance/features/products"
)

type productQuery struct {
	db *sql.DB
}

func New(db *sql.DB) products.Data {
	return &productQuery{
		db: db,
	}
}

func (p *productQuery) Get(data *products.ProductsCode) (*[]products.ProductsCode, error) {
	panic("unimplemented")
}

func (p *productQuery) GetById(productId string) (*products.ProductsCode, error) {
	panic("unimplemented")
}

func (p *productQuery) CheckAvailabitity(data *products.ProductsCode) (*products.ProductsCode, error) {
	panic("unimplemented")
}
