package helpers

import (
	"golang.org/x/crypto/bcrypt"
)

func Encrypt(str string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(str), bcrypt.DefaultCost)
	return string(hash), err
}

func CompareEncrypt(str, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(str))
	return err == nil
}
