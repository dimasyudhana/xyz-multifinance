package factory

import (
	"database/sql"

	_productData "github.com/dimasyudhana/xyz-multifinance/features/products/data"
	_productHandler "github.com/dimasyudhana/xyz-multifinance/features/products/handler"
	_productService "github.com/dimasyudhana/xyz-multifinance/features/products/service"
)

func New(db *sql.DB) *_productHandler.ProductHandler {
	data := _productData.New(db)
	service := _productService.New(data)
	handler := _productHandler.New(service)
	return handler
}
