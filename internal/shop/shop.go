package shop

import (
	"errors"

	"github.com/theo-krutiy/minimal-go/internal/models"
)

type Database interface {
	ReadItems(query string, offset, limit int) (page []*models.Item, totalResults int, err error)
	GetCart(userId string) (*models.Cart, error)
	IncrementItemInCart(cartId, itemId string, increment int) error
	EmptyCart(cartId string) error
	ConvertCartIntoOrder(cartId string) (string, error)
}

func GetItems(query string, offset, limit int, db Database) (page []*models.Item, totalResults int, err error) {
	if limit <= 0 {
		err = errors.New("limit must be positive")
		return
	}
	if offset < 0 {
		err = errors.New("offset must not be negative")
		return
	}

	page, totalResults, err = db.ReadItems(query, offset, limit)
	return
}

func GetCart(userId string, db Database) (cart *models.Cart, err error) {
	cart, err = db.GetCart(userId)
	if err != nil {
		return nil, err
	}
	return cart, nil
}

func AddItemToCart(cartId, itemId string, countAdded *int, db Database) error {
	switch {
	case *countAdded <= 0:
		return errors.New("count must be positive")
	case countAdded == nil:
		v := 1
		countAdded = &v
	}

	err := db.IncrementItemInCart(cartId, itemId, *countAdded)
	return err
}

func RemoveItemFromCart(cartId, itemId string, countRemoved *int, db Database) error {
	switch {
	case countRemoved == nil:
		v := 1
		countRemoved = &v
	case *countRemoved <= 0:
		return errors.New("count must be positive")
	}

	err := db.IncrementItemInCart(cartId, itemId, *countRemoved*-1)
	return err
}

func EmptyCart(cartId string, db Database) error {
	err := db.EmptyCart(cartId)
	return err
}

func ConvertCartIntoOrder(cartId string, db Database) (string, error) {
	orderId, err := db.ConvertCartIntoOrder(cartId)
	if err != nil {
		return "", err
	}

	return orderId, nil
}
