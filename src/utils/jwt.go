package utils

import (
	"github.com/golang-jwt/jwt"
	"time"
)

func CreateJWT(phoneNumber string) (string, error) {
	mySigningKey := []byte("example")

	aToken := jwt.New(jwt.SigningMethodHS256)
	claims := aToken.Claims.(jwt.MapClaims)
	claims["PhoneNum"] = phoneNumber
	claims["exp"] = time.Now().Add(time.Minute * 20).Unix()

	token, err := aToken.SignedString(mySigningKey)
	if err != nil {
		return "", err
	}
	return token, nil
}
