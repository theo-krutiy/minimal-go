package auth

import (
	"errors"

	"github.com/theo-krutiy/minimal-go/internal/models"
)

type db interface {
	CreateNewUser(login, passwordHash string) (string, error)
	ReadUser(user *models.UserInDatabase) error
}

func CreateNewUser(login, password string, db db) (string, error) {
	pwdHash, err := hashPassword(password)
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

func ValidateCredentials(login, password string, db db) (string, error) {
	user := models.UserInDatabase{Login: login}
	if err := db.ReadUser(&user); err != nil {
		return "", errors.New("couldn't read user in database")
	}
	if err := comparePassword(password, user.PasswordHash); err != nil {
		return "", errors.New("incorrect password")
	}

	token, err := issueNewJWT()
	if err != nil {
		return "", errors.New("error issuing jwt")
	}

	return token, nil
}

func ValidateJWT(token string) error {
	return validateJWT(token)
}
