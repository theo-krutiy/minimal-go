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

type ItemInCart struct {
	ItemId string `json:"item_id" db:"item_id"`
	CartId string `json:"cart_id" db:"cart_id"`
	Count  int    `json:"count" db:"count_in_cart"`
}

type Cart struct {
	Id     string        `json:"id" db:"id"`
	UserId string        `json:"user_id" db:"user_id"`
	Items  []*ItemInCart `json:"items"`
}
