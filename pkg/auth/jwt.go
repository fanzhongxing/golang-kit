package auth

import (
	"errors"
	"github.com/golang-jwt/jwt/v5"
	"time"
)

// CreateToken 生成 jwt token
func CreateToken(data map[string]interface{}, secret string) (string, error) {
	claims := make(jwt.MapClaims)
	claims = data
	claims["iat"] = time.Now().Unix()
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(secret))
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

// DecodeToken token 解码
func DecodeToken(tokenStr string, secret string) (jwt.MapClaims, error) {
	token, err := jwt.ParseWithClaims(tokenStr, jwt.MapClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(secret), nil
	})
	if err != nil {
		return nil, err
	}
	tokenClaims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return nil, errors.New("invalid type")
	}
	return tokenClaims, nil
}
