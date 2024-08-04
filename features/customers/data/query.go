package data

import (
	"database/sql"
	"errors"

	"github.com/dimasyudhana/xyz-multifinance/features/customers"
	"github.com/google/uuid"
)

type customerQuery struct {
	db *sql.DB
}

func New(db *sql.DB) customers.Data {
	return &customerQuery{
		db: db,
	}
}

func (uq *customerQuery) Insert(data *customers.CustomersCore) (*customers.CustomersCore, error) {
	queryData := CoreToData(data)
	queryData.CustomerId = uuid.New().String()

	stmt, err := uq.db.Prepare(
		"INSERT INTO users (customer_id, email, phone, password, nik, fullname, legalname, pob, dob, salary, ocr_picture, liveness_picture, `limit`, created_at, updated_at) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, NOW(), NOW());",
	)
	if err != nil {
		return nil, errors.New("failed to prepare statement: " + err.Error())
	}
	defer stmt.Close()

	_, err = stmt.Exec(
		queryData.CustomerId,
		queryData.Email,
		queryData.Phone,
		queryData.Password,
		queryData.Nik,
		queryData.Fullname,
		queryData.Legalname,
		queryData.Pob,
		queryData.Dob,
		queryData.Salary,
		queryData.OCRPicture,
		queryData.LivenessPicture,
		queryData.Limit,
	)
	if err != nil {
		return nil, errors.New("failed to execute statement: " + err.Error())
	}

	return DataToCore(queryData), nil
}
