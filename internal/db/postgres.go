package db

import (
	"context"
	"errors"

	"github.com/jackc/pgx"
	"github.com/jackc/pgx/v5/pgxpool"

	"github.com/theo-krutiy/minimal-go/internal/models"
)

type Postgres struct {
	pool *pgxpool.Pool
}

func NewPostgres(ctx context.Context, connString string) (*Postgres, error) {
	pool, err := pgxpool.New(ctx, connString)
	if err != nil {
		return nil, err
	}
	if err = pool.Ping(ctx); err != nil {
		return nil, err
	}
	return &Postgres{pool}, nil
}

func (p *Postgres) CreateNewUser(login, passwordHash string) (string, error) {
	var userId string
	err := p.pool.QueryRow(context.Background(), "INSERT INTO users (login, password_hash) VALUES ($1, $2)  RETURNING id;", login, passwordHash).Scan(&userId)
	if err != nil {
		return "", err
	}
	return userId, nil
}

func (p *Postgres) ReadUser(user *models.UserInDatabase) error {
	err := p.pool.QueryRow(context.Background(), "SELECT id, password_hash FROM users WHERE login = $1;", user.Login).Scan(&user.Id, &user.PasswordHash)
	switch err {
	case nil:
		return nil
	case pgx.ErrNoRows:
		return errors.New("unknown login")
	default:
		return errors.New("unknown error")
	}
}
