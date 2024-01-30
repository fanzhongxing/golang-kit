package auth

import (
	"github.com/golang-jwt/jwt/v5"
	"time"
)

var _DefaultSecret = "cbt0121"

// CreateToken 生成 jwt token
func CreateToken(data map[string]interface{}) (string, error) {
	claims := make(jwt.MapClaims)
	claims = data
	claims["iat"] = time.Now().Unix()
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(_DefaultSecret))
	if err != nil {
		return "", err
	}
	return tokenString, nil
}
