package authentication

type CustomerCore struct {
	CustomerId string
	Email      string `validate:"required,email"`
	Password   string `validate:"required"`
	Role       int
}

type AuthDataInterface interface {
	GetCustomerByEmail(email string) (*CustomerCore, error)
}

type AuthServiceInterface interface {
	Login(email string, password string) (*CustomerCore, string, error)
}
