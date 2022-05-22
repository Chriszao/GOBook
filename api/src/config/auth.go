package config

import (
	"errors"
	"fmt"
	"net/http"
	"strings"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
)

func GenerateToken(userId uint64) (string, error) {
	claims := jwt.MapClaims{}

	claims["authorized"] = true
	claims["exp"] = time.Now().Add(time.Hour * 6).Unix()
	claims["userId"] = userId

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString(JwtSecretKey)
}

// Verifies if the token is valid
func ValidateToken(request *http.Request) error {
	tokenString := getToken(request)

	token, err := jwt.Parse(tokenString, keyFunc)

	if err != nil {
		return err
	}

	if _, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return nil
	}

	return errors.New("invalid token")
}

func getToken(request *http.Request) string {
	token := request.Header.Get("Authorization")

	if len(strings.Split(token, " ")) == 2 {
		return strings.Split(token, " ")[1]
	}

	return ""
}

func keyFunc(token *jwt.Token) (interface{}, error) {
	if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
		return nil, fmt.Errorf("unexpected signature method %v", token.Header["alg"])
	}

	return JwtSecretKey, nil
}
