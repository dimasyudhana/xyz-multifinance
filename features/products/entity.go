package products

import "time"

type ProductsCode struct {
	ProductId string     `json:"product_id"`
	Category  string     `json:"category"`
	SKU       string     `json:"sku"`
	Name      string     `json:"name"`
	Amount    float32    `json:"amount"`
	Stock     uint32     `json:"stock"`
	Tenor     uint32     `json:"tenor"`
	Bunga     uint32     `json:"bunga"`
	Fee       uint32     `json:"fee"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
	DeletedAt *time.Time `json:"deleted_at"`
}

type Service interface {
	Get(data *ProductsCode) (*[]ProductsCode, error)
	GetById(productId string) (*ProductsCode, error)
	CheckAvailabitity(data *ProductsCode) (*ProductsCode, error)
}

type Data interface {
	Get(data *ProductsCode) (*[]ProductsCode, error)
	GetById(productId string) (*ProductsCode, error)
	CheckAvailabitity(data *ProductsCode) (*ProductsCode, error)
}
