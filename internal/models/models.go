package models

type User struct {
	Id           string `db:"id"`
	Login        string `db:"login"`
	PasswordHash []byte `db:"password_hash"`
}

type Item struct {
	Id             string `json:"id" db:"id"`
	Name           string `json:"name" db:"name"`
	CountAvailable int    `json:"count_available" db:"count_available"`
	PriceInteger   int    `json:"price_integer" db:"price_integer"`
	PriceDecimal   int    `json:"price_decimal" db:"price_decimal"`
}
