package common

import (
	"golang.org/x/crypto/bcrypt"
)

func PasswordHash(plainText string) (string, error) {
	h, err := bcrypt.GenerateFromPassword([]byte(plainText), bcrypt.DefaultCost)

	return string(h), err
}

func CheckPassword(plainText string, hashedText string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedText), []byte(plainText))
}
