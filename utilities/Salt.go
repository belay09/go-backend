package utilities

import (
	"golang.org/x/crypto/bcrypt"
)


// var salt = []byte{15, 156, 78, 209, 26, 116, 116, 20, 126, 226, 123, 212, 42, 148, 117, 207}
func HashPassword(password string) ([]byte, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), 8)
	if err != nil {
		return nil, err
	}
	return hashedPassword, nil
}

func CompareHashAndPassword(dbPassword, userPassword string) bool {
	return bcrypt.CompareHashAndPassword([]byte(dbPassword), []byte(userPassword)) == nil
}
