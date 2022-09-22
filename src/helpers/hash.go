package helpers

import "golang.org/x/crypto/bcrypt"

// Register
func HashPassword(pass string) (string, error) {
	hashPass, err := bcrypt.GenerateFromPassword([]byte(pass), bcrypt.DefaultCost)

	if err != nil {
		return "", err
	}

	return string(hashPass), nil
}

// Login
func CheckPassword(hashedPassword, passwordDb string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(passwordDb))
	return err == nil
}
