package utils

import (
	"example.com/m/src/domain"
	"github.com/golang-jwt/jwt"
	"time"
)

var	mySigningKey = []byte("example")

func CreateJWT(phoneNumber string, userId int) (string, error) {

	aToken := jwt.New(jwt.SigningMethodHS256)
	claims := aToken.Claims.(jwt.MapClaims)
	claims["PhoneNum"] = phoneNumber
	claims["UserId"] = userId
	claims["exp"] = time.Now().Add(time.Minute * 20).Unix()

	token, err := aToken.SignedString(mySigningKey)
	if err != nil {
		return "", err
	}
	return token, nil
}

func VerifyToken(values string) (int, error) {
	if len(values) == 0 {
		return 0, domain.ErrRequiredAccessToken
	}

	token, err := jwt.Parse(values, func(token *jwt.Token) (interface{}, error) {
		return mySigningKey, nil
	})
	if err != nil {
		switch err.(*jwt.ValidationError).Errors {
		case jwt.ValidationErrorExpired:
			return 0, domain.ErrExpiredToken
		default:
			return 0, domain.ErrInvalidAccessToken
		}
	}

	var value int
	if claims, ok := token.Claims.(jwt.MapClaims); ok {
		if userIdFloat, ok := claims["UserId"].(float64); ok {
			value = int(userIdFloat)
		}
	}
	return value, nil
}
