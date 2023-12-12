package auth

import (
	"golang.org/x/crypto/bcrypt"
)

func HashPassword(password []byte) ([]byte, error) {
	return hashPassword(password)
}

func hashPassword(password []byte) ([]byte, error) {
	return bcrypt.GenerateFromPassword(password, bcrypt.DefaultCost)
}

func VerifyHash(hash []byte) error {
	return verifyHash(hash)
}

func verifyHash(hash []byte) error {
	return bcrypt.CompareHashAndPassword(hash, []byte("Lohith@12"))
}
