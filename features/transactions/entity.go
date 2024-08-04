package transactions

type TransactionCore struct {
	TransactionId   string `json:"transaction_id"`
	CustomerId      string `json:"customer_id"`
	NomorKontrak    string `json:"nomor_kontrak"`
	NamaAsset       string `json:"nama_asset"`
	OTR             string `json:"otr"`
	JumlahCicilan   string `json:"jumlah_cicilan"`
	JumlahBunga     string `json:"jumlah_bunga"`
	AdminFee        string `json:"admin_fee"`
	TransactionDate string `json:"transaction_date"`
}

type TransactionProductCore struct {
	CustomerId string `json:"customer_id"`
	ProductId  string `json:"product_id"`
}

type Data interface {
	Create(data *TransactionProductCore) (*TransactionCore, error)
	CheckStatus(data *TransactionProductCore) (*TransactionCore, error)
	GetById(transactionId string) (*TransactionCore, error)
}

type Service interface {
	Create(data *TransactionProductCore) (*TransactionCore, error)
	CheckStatus(data *TransactionProductCore) (*TransactionCore, error)
	GetById(transactionId string) (*TransactionCore, error)
}
