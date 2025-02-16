package db

import (
	"context"
	"errors"
	"fmt"

	"github.com/jackc/pgx/v5"
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

func (p *Postgres) CreateNewUser(login string, passwordHash []byte) (string, error) {
	var userId string
	err := p.pool.QueryRow(context.Background(), "INSERT INTO users (login, password_hash) VALUES ($1, $2)  RETURNING id;", login, passwordHash).Scan(&userId)
	if err != nil {
		return "", err
	}
	return userId, nil
}

func (p *Postgres) ReadUser(user *models.User) error {
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

func (p *Postgres) ReadItems(query string, offset, limit int) (page []*models.Item, totalResults int, err error) {
	query = fmt.Sprintf("%v%%", query)
	batch := &pgx.Batch{}
	batch.Queue(
		"SELECT * FROM items WHERE name LIKE $1 ORDER BY name OFFSET $2 LIMIT $3;",
		query, offset, limit,
	).Query(func(rows pgx.Rows) error {
		page, err = pgx.CollectRows(rows, pgx.RowToAddrOfStructByNameLax[models.Item])
		return err
	})
	batch.Queue("SELECT COUNT(*) FROM items WHERE name LIKE $1;", query).QueryRow(func(row pgx.Row) error {
		err := row.Scan(&totalResults)
		return err
	})

	err = p.pool.SendBatch(context.Background(), batch).Close()
	return
}
