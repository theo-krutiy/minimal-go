package db

type Database interface {
	CreateNewUser(login, passwordHash string) (string, error)
}
