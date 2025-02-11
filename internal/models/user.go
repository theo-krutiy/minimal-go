package models

type UserInDatabase struct {
	Id           string
	Login        string
	PasswordHash []byte
}
