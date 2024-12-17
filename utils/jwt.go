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

func VerifyToken(tokenString string) (string, int, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(private_key), nil
	})

	if err != nil {
		return "", 0, err
	}
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		fmt.Println("CLAIMS", claims)
		email := claims["email"].(string)
		userId := int(claims["userId"].(float64))
		return email, userId, nil
	}

	return "", 0, fmt.Errorf("invalid token")
}
