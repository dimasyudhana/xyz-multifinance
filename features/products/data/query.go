package data

import (
	"database/sql"
	"errors"

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

func (p *productQuery) Get(keyword string) ([]*products.ProductsCore, error) {

	search := "%" + keyword + "%"

	stmt, err := p.db.Prepare(`
		SELECT product_id, category, sku, name, amount, stock, tenor, bunga, fee
		FROM products where name like ?;
	`)
	if err != nil {
		return nil, errors.New("failed to prepare statement: " + err.Error())
	}
	defer stmt.Close()

	rows, err := stmt.Query(search)
	if err != nil {
		return nil, errors.New("failed to execute statement: " + err.Error())
	}
	defer rows.Close()

	var product []*products.ProductsCore
	for rows.Next() {
		var pr products.ProductsCore
		err := rows.Scan(
			&pr.ProductId,
			&pr.Category,
			&pr.SKU,
			&pr.Name,
			&pr.Amount,
			&pr.Stock,
			&pr.Tenor,
			&pr.Bunga,
			&pr.Fee,
		)
		if err != nil {
			return nil, errors.New("failed to scan row: " + err.Error())
		}
		product = append(product, &pr)
	}

	if err = rows.Err(); err != nil {
		return nil, errors.New("error occurred while iterating rows: " + err.Error())
	}

	return product, nil
}

func (p *productQuery) GetById(productId string) (*products.ProductsCore, error) {
	panic("unimplemented")
}

func (p *productQuery) CheckAvailabitity(data *products.ProductsCore) (*products.ProductsCore, error) {
	panic("unimplemented")
}
