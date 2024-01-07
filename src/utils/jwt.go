package utils

import (
	"example.com/m/src/domain"
	"github.com/golang-jwt/jwt"
	goJwt "github.com/golang-jwt/jwt/v5"
	"strings"
	"time"
)

func CreateJWT(phoneNumber string, userId int) (string, error) {
	mySigningKey := []byte("example")

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

func GetClaimByUserId(values string) (int, error) {
	if len(values) == 0 {
		return 0, domain.ErrRequiredAccessToken
	}

	accessToken := RemoveBearer(values)
	token, _, err := new(goJwt.Parser).ParseUnverified(accessToken, goJwt.MapClaims{})
	if err != nil {
		return 0, domain.ErrInvalidAccessToken
	}

	var value int
	if claims, ok := token.Claims.(goJwt.MapClaims); ok {
		if userIDFloat, ok := claims["UserId"].(float64); ok {
			value = int(userIDFloat)
		}
	}
	return value, nil
}

// RemoveBearer token 에서 Bearer 제거
func RemoveBearer(accessToken string) string {
	removeBearer := strings.ReplaceAll(accessToken, "Bearer ", "")
	removeBearer = strings.ReplaceAll(removeBearer, "bearer ", "")
	return removeBearer
}
