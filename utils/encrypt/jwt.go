package encrypt

import (
	"fmt"
	"time"

	"github.com/dimasyudhana/xyz-multifinance/app/config"
	"github.com/golang-jwt/jwt/v5"
)

func GenerateToken(userId string) (string, interface{}) {
	claims := jwt.MapClaims{}
	claims["authorized"] = true
	claims["userId"] = userId
	claims["exp"] = time.Now().Add(time.Hour * 3).Unix()
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(config.SECRET_JWT))
}

func ExtractToken(t interface{}) (userId any, role string, err error) {
	user := t.(*jwt.Token)
	if user.Valid {
		claims := user.Claims.(jwt.MapClaims)
		switch claims["userId"].(type) {
		case float64:
			userId = claims["userId"].(float64)
		case string:
			userId = claims["userId"].(string)
		default:
			return 0, "", fmt.Errorf("token invalid")
		}
		role := claims["role"].(string)
		return userId, role, nil
	}
	return 0, "", fmt.Errorf("token invalid")
}
