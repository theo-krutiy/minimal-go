package auth

import "golang.org/x/crypto/bcrypt"

func hashPassword(pwd []byte) ([]byte, error) {
	hash, err := bcrypt.GenerateFromPassword(pwd, bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}
	return hash, nil
}

func comparePassword(pwd, hash []byte) error {
	err := bcrypt.CompareHashAndPassword(hash, pwd)
	return err
}
