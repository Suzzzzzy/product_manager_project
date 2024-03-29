package utils

import "golang.org/x/crypto/bcrypt"

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(bytes), err
}

func CheckPasswordHash(hashVal, password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashVal), []byte(password))
	if err != nil {
		return false
	} else {
		return true
	}
}
