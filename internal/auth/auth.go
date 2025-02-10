package auth

import (
	"errors"
)

type db interface {
	CreateNewUser(login, passwordHash string) (string, error)
	ReadUser(user *userInDatabase) error
}

type userInDatabase struct {
	Id           string
	Login        string
	PasswordHash string
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
	user := userInDatabase{Login: login}
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
