package utils

import "golang.org/x/crypto/bcrypt"

func HashPassword(password string) (string, error) {
	// 14 defines the difficulty level of password generated
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)

	return string(bytes), err
}

func CheckPasswordHash(password, hashedPassword string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	return err == nil
}
