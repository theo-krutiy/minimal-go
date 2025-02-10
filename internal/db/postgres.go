package db

type Postgres struct {
}

func NewPostgres() *Postgres {
	return &Postgres{}
}

func (p *Postgres) CreateNewUser(login, passwordHash string) (string, error) {
	return "dummyIdFromDB", nil
}
