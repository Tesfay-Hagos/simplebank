package util

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

func HashedPassword(password string) (string, error) {
	hashpassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)

	if err != nil {
		return "", fmt.Errorf("hashing failed:%s", err)
	}
	return string(hashpassword), nil
}

func CheckHashedpassword(password, hasehedpassword string) error {
	err := bcrypt.CompareHashAndPassword([]byte(hasehedpassword), []byte(password))
	if err != nil {
		return err
	}
	return nil
}
