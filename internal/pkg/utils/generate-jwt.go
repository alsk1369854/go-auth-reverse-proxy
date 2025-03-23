package utils

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
)

func GenerateJWT(name string, secretKey string) (string, error) {
	claims := jwt.MapClaims{
		"name": name,
		"iat": time.Now().Unix(),                   // 發行時間
		// "exp":      time.Now().Add(time.Hour * 1).Unix(), // 過期時間
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(secretKey))
}