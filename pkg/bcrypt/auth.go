package bcrypt

import (
	"fmt"
	"log"

	"golang.org/x/crypto/bcrypt"
)

func EncryptPassword(password string) (string, error) {
	encrypted, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)

	if err != nil {
		log.Fatal(err)
		return "", fmt.Errorf("bycrypt: failed to encrypt password")
	}

	return string(encrypted), nil
}

func CompareHashAndPassword(encrypted string, password string) bool {

	return bcrypt.CompareHashAndPassword([]byte(encrypted), []byte(password)) == nil
}
