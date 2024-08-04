package data

import (
	"database/sql"
	"errors"

	"github.com/dimasyudhana/xyz-multifinance/features/authentication"
)

type authQuery struct {
	db *sql.DB
}

func New(db *sql.DB) authentication.AuthDataInterface {
	return &authQuery{
		db: db,
	}
}

func (query *authQuery) Create(data *authentication.CustomerCore) (string, error) {
	queryData := CoreToData(data)
	stmt, err := query.db.Prepare(
		"INSERT INTO customers (customer_id, email, phone, password, fullname, legalname, nik, pob, dob, salary, ocr_picture, liveness_picture, `limit`, created_at, updated_at) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, NOW(), NOW());",
	)
	if err != nil {
		return "", errors.New("failed to prepare statement: " + err.Error())
	}
	defer stmt.Close()
	_, err = stmt.Exec(
		queryData.CustomerId,
		queryData.Email,
		queryData.Phone,
		queryData.Password,
		queryData.Fullname,
		queryData.Legalname,
		queryData.Nik,
		queryData.Pob,
		queryData.Dob,
		queryData.Salary,
		queryData.OCRPicture,
		queryData.LivenessPicture,
		queryData.Limit,
	)
	if err != nil {
		return "", errors.New("failed to execute statement: " + err.Error())
	}
	return queryData.CustomerId, nil
}

func (query *authQuery) GetCustomerByEmail(email string) (*authentication.CustomerCore, error) {
	var auth authentication.CustomerCore
	if err := query.db.QueryRow("SELECT customer_id, password, email FROM customers WHERE email = ? AND deleted_at IS NULL;", email).Scan(&auth.CustomerId, &auth.Password, &auth.Email); err != nil {
		return nil, err
	}
	return &auth, nil
}
