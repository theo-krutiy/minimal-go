package auth

import (
	"errors"

	"github.com/theo-krutiy/minimal-go/internal/models"
)

type Database interface {
	CreateNewUser(login string, passwordHash []byte) (string, error)
	ReadUser(user *models.User) error
}

func CreateNewUser(login, password string, db Database) (string, error) {
	pwdBytes := []byte(password)
	if len(pwdBytes) > 72 {
		return "", errors.New("password too long")
	}
	pwdHash, err := hashPassword(pwdBytes)
	if err != nil {
		return "", errors.New("couldn't hash password")
	}

	// create user in db, return id
	dummyId, err := db.CreateNewUser(login, pwdHash)
	if err != nil {
		return "", errors.New("couldn't create new user in database")
	}
	return dummyId, nil
}

func Authenticate(login, password string, secret []byte, db Database) (string, error) {
	if err := ValidateCredentials(login, password, db); err != nil {
		return "", err
	}

	token, err := issueNewJWT(secret)
	if err != nil {
		return "", errors.New("error issuing jwt")
	}

	return token, nil
}

func ValidateCredentials(login, password string, db Database) error {
	pwdBytes := []byte(password)
	if len(pwdBytes) > 72 {
		return errors.New("incorrect password")
	}
	user := models.User{Login: login}
	if err := db.ReadUser(&user); err != nil {
		return errors.New("couldn't read user in database")
	}
	if err := comparePassword(pwdBytes, user.PasswordHash); err != nil {
		return errors.New("incorrect password")
	}
	return nil
}

func ValidateToken(token string, secret []byte) error {
	return validateJWT(token, secret)
}
