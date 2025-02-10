package db

import "github.com/theo-krutiy/minimal-go/internal/models"

type Database interface {
	CreateNewUser(login, passwordHash string) (string, error)
	ReadUser(user *models.UserInDatabase) error
}
