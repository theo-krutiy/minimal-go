package shop

import (
	"errors"

	"github.com/theo-krutiy/minimal-go/internal/models"
)

type Database interface {
	ReadItems(query string, offset, limit int) (page []*models.ItemInDatabase, totalResults int, err error)
}

func GetItems(query string, offset, limit int, db Database) (page []*models.ItemInDatabase, totalResults int, err error) {
	if limit == 0 {
		err = errors.New("limit 0")
		return
	}

	page, totalResults, err = db.ReadItems(query, offset, limit)
	return
}
