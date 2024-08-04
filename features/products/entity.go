package products

type ProductsCore struct {
	ProductId string `json:"product_id"`
	Category  string `json:"category"`
	SKU       string `json:"sku"`
	Name      string `json:"name"`
	Amount    string `json:"amount"`
	Stock     string `json:"stock"`
	Tenor     string `json:"tenor"`
	Bunga     string `json:"bunga"`
	Fee       string `json:"fee"`
	CreatedAt string `json:"created_at,omitempty"`
	UpdatedAt string `json:"updated_at,omitempty"`
	DeletedAt string `json:"deleted_at,omitempty"`
}

type Service interface {
	Get(keyword string) (<-chan *ProductsCore, error)
	GetById(productId string) (*ProductsCore, error)
	CheckAvailabitity(data *ProductsCore) (*ProductsCore, error)
}

type Data interface {
	Get(keyword string) ([]*ProductsCore, error)
	GetById(productId string) (*ProductsCore, error)
	CheckAvailabitity(data *ProductsCore) (*ProductsCore, error)
}
