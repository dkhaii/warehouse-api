package helpers

import "golang.org/x/crypto/bcrypt"

func HashPassword(password string) (string, error) {
	passwordBytes := []byte(password)

	hashedPassword, err := bcrypt.GenerateFromPassword(passwordBytes, bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}

	return string(hashedPassword), err
}

func ComparePassword(hashedPassword string, plainPassword string) bool {
	plainPasswordBytes := []byte(plainPassword)
	hashedPasswordBytes := []byte(hashedPassword)

	err := bcrypt.CompareHashAndPassword(hashedPasswordBytes, plainPasswordBytes)

	return err == nil
}
