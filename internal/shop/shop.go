package shop

import (
	"errors"
	"fmt"

	"github.com/theo-krutiy/minimal-go/internal/errcodes"
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
		err = fmt.Errorf("limit must be positive%w", errcodes.Validation)
		return
	}
	if offset < 0 {
		err = fmt.Errorf("offset must not be negative%w", errcodes.Validation)
		return
	}

	page, totalResults, err = db.ReadItems(query, offset, limit)
	if err != nil {
		err = errcodes.Unknown
	}
	return
}

func GetCart(userId string, db Database) (cart *models.Cart, err error) {
	cart, err = db.GetCart(userId)
	switch {
	case err == nil:
		return cart, nil
	case errors.Is(err, errcodes.DBNoData):
		return nil, errcodes.NoData
	default:
		return nil, errcodes.Unknown
	}
}

func AddItemToCart(cartId, itemId string, countAdded *int, db Database) error {
	switch {
	case *countAdded <= 0:
		return fmt.Errorf("count must be positive%w", errcodes.Validation)
	case countAdded == nil:
		v := 1
		countAdded = &v
	}

	err := db.IncrementItemInCart(cartId, itemId, *countAdded)
	if err != nil {
		return errcodes.Unknown
	}
	return nil
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
	if err != nil {
		return errcodes.Unknown
	}
	return nil
}

func EmptyCart(cartId string, db Database) error {
	err := db.EmptyCart(cartId)
	if err != nil {
		return errcodes.Unknown
	}
	return nil
}

func ConvertCartIntoOrder(cartId string, db Database) (string, error) {
	orderId, err := db.ConvertCartIntoOrder(cartId)
	if err != nil {
		return "", errcodes.Unknown
	}

	return orderId, nil
}
