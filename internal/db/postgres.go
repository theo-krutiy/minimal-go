package db

import "github.com/theo-krutiy/minimal-go/internal/models"

type Postgres struct {
}

func NewPostgres() *Postgres {
	return &Postgres{}
}

func (p *Postgres) CreateNewUser(login, passwordHash string) (string, error) {
	return "dummyIdFromDB", nil
}

func (p *Postgres) ReadUser(user *models.UserInDatabase) error {
	user.Id = "dummyIdFromDB"
	user.PasswordHash = "dummyPasswordHashFromDB"

	return nil
}
