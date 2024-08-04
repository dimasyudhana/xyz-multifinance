package data

import (
	"database/sql"
	"errors"

	"github.com/dimasyudhana/xyz-multifinance/features/customers"
)

type customerQuery struct {
	db *sql.DB
}

func New(db *sql.DB) customers.Data {
	return &customerQuery{
		db: db,
	}
}

func (uq *customerQuery) CustomerTransactions(customerId string) ([]*customers.CustomersCore, error) {
	stmt, err := uq.db.Prepare(`
		SELECT 
			c.customer_id, c.email, c.phone, c.fullname, 
			t.nomor_kontrak, t.nama_asset, t.otr, t.admin_fee, t.jumlah_cicilan, t.jumlah_bunga, t.transaction_date
		FROM customers c
		JOIN transactions t ON t.customer_id = c.customer_id
		WHERE c.customer_id = ?;
	`)
	if err != nil {
		return nil, errors.New("failed to prepare statement: " + err.Error())
	}
	defer stmt.Close()

	rows, err := stmt.Query(customerId)
	if err != nil {
		return nil, errors.New("failed to execute statement: " + err.Error())
	}
	defer rows.Close()

	var customer []*customers.CustomersCore
	for rows.Next() {
		var ct customers.CustomersCore
		err := rows.Scan(
			&ct.CustomerId,
			&ct.Email,
			&ct.Phone,
			&ct.Fullname,
			&ct.NomorKontrak,
			&ct.NamaAsset,
			&ct.OTR,
			&ct.AdminFee,
			&ct.JumlahCicilan,
			&ct.JumlahBunga,
			&ct.TransactionDate,
		)
		if err != nil {
			return nil, errors.New("failed to scan row: " + err.Error())
		}
		customer = append(customer, &ct)
	}

	if err = rows.Err(); err != nil {
		return nil, errors.New("error occurred while iterating rows: " + err.Error())
	}

	return customer, nil
}
