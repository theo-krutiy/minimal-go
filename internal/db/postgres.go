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

func (p *Postgres) GetCart(userId string) (*models.Cart, error) {
	cart := &models.Cart{UserId: userId}
	err := p.pool.QueryRow(context.Background(), "SELECT id FROM carts WHERE user_id = $1;", userId).Scan(&cart.Id)
	switch {
	case errors.Is(err, pgx.ErrNoRows):
		return nil, errors.New("cart doesn't exist")
	case err != nil:
		return nil, errors.New("unknown error")
	}
	rows, err := p.pool.Query(context.Background(), "SELECT item_id, count_in_cart FROM items_in_cart WHERE cart_id = $1;", cart.Id)
	if err != nil {
		return nil, err
	}
	cart.Items, err = pgx.CollectRows(rows, pgx.RowToAddrOfStructByNameLax[models.ItemInCart])
	if err != nil {
		return nil, err
	}
	return cart, nil
}

func (p *Postgres) CreateCart(userId string) (string, error) {
	var cartId string
	err := p.pool.QueryRow(context.Background(), "INSERT INTO carts (user_id) VALUES $1 RETURNING id;", userId).Scan(&cartId)
	if err != nil {
		return "", err
	}

	return cartId, nil
}

// TODO: handle negative increment
func (p *Postgres) IncrementItemInCart(cartId, itemId string, increment int) error {
	query := `
		INSERT INTO items_in_cart (cart_id, item_id, count_in_cart)
		VALUES ($1, $2, $3)
		ON CONFLICT (cart_id, item_id) DO UPDATE 
		SET count_in_cart = items_in_cart.count_in_cart + EXCLUDED.count_in_cart
	`

	_, err := p.pool.Exec(context.Background(), query, cartId, itemId, increment)
	return err
}

func (p *Postgres) EmptyCart(cartId string) error {
	query := `DELETE FROM items_in_cart WHERE cart_id = $1`
	_, err := p.pool.Exec(context.Background(), query, cartId)
	return err
}

func (p *Postgres) ConvertCartIntoOrder(cartId string) (string, error) {
	return "", nil
}
