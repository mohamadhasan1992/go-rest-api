package utils

import "golang.org/x/crypto/bcrypt"

var salt = 14

func HashPassword(password string) (string, error) {
	hashedBytes, err := bcrypt.GenerateFromPassword([]byte(password), salt)
	return string(hashedBytes), err
}

func CompareHash(password, hashedPassword string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	return err == nil
}
