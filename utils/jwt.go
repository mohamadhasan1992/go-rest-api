package utils

import (
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

const secretKey = "this_is_a_secret_jwt"

func GenerateToken(email string, userId int64) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"email":  email,
		"userId": userId,
		"exp":    time.Now().Add(time.Hour * 12).Unix(),
	})
	return token.SignedString([]byte(secretKey))
}

func VerifyToken(token string) (int64, error) {
	token = strings.TrimPrefix(token, "Bearer ")
	parsedToken, err := jwt.Parse(token, func(token *jwt.Token) (any, error) {
		_, ok := token.Method.(*jwt.SigningMethodHMAC)
		if !ok {
			fmt.Println("Unexpected signing method")
			return nil, errors.New("Unexpected signing method")
		}
		return []byte(secretKey), nil
	})

	if err != nil {
		return 0, errors.New("Could not parse token")
	}
	tokenIsValid := parsedToken.Valid
	if !tokenIsValid {
		fmt.Println("Invalid Token!")
		return 0, errors.New("Invalid Token!")
	}
	claims, ok := parsedToken.Claims.(jwt.MapClaims)
	if !ok {
		return 0, errors.New("Invalid Token!")
	}
	// email := claims["email"].(string)
	userId := claims["userId"].(float64)
	return int64(userId), nil
}
