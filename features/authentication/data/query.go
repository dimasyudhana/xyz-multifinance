package data

import (
	"database/sql"

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

func (query *authQuery) GetCustomerByEmail(email string) (*authentication.CustomerCore, error) {
	var auth authentication.CustomerCore
	if err := query.db.QueryRow("SELECT user_id, role, password, email FROM xyz_multifinance.users WHERE email = ? AND deleted_at IS NULL;", email).Scan(&auth.CustomerId, &auth.Role, &auth.Password, &auth.Email); err != nil {
		return nil, err
	}
	return &auth, nil
}
