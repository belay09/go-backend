package utilities

import (
	"crypto/rand"
	"golang.org/x/crypto/bcrypt"
)

func GenerateSalt() []byte {
	salt := make([]byte, 16)
	_, err := rand.Read(salt)
	if err != nil {
		panic(err) 
	}
	return salt
}

func HashPassword(password string, salt []byte) ([]byte, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword(append(salt, []byte(password)...), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}
	return hashedPassword, nil
}