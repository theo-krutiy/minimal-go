package auth

import "golang.org/x/crypto/bcrypt"

func hashPassword(pwd string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(pwd), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hash), nil
}

func comparePassword(pwd, hash string) error {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(pwd))
	return err
}
