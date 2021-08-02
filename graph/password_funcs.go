package graph

import (
	"golang.org/x/crypto/bcrypt"
)

// HashPassword password to bytes
func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	if err != nil {
		Logger(err)
	}
	return string(bytes), err
}

// CheckPasswordHash compare hashed passwords
func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
