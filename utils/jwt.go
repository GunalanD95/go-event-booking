package utils

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

const private_key = "3787r83undjnudfbdjcjkbfdfjkdsjjsj@f"

func GenerateTokenJwt(username string, userId int) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"email":  username,
		"userId": userId,
		"exp":    time.Now().Add(time.Hour * 2).Unix(),
	})
	return token.SignedString([]byte(private_key))
}

func VerifyToken(tokenString string) (bool, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(private_key), nil
	})
	if err != nil {
		return false, err
	}
	if token.Valid {
		return true, nil
	}
	return false, fmt.Errorf("invalid token")
}
