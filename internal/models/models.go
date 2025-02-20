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
	Price          int    `json:"price" db:"price"`
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

type Order struct {
	Id          string         `json:"id" db:"id"`
	UserId      string         `json:"user_id" db:"user_id"`
	TotalPrice  int            `json:"total_price" db:"total_price"`
	State       string         `json:"state" db:"state"`
	CreatedAt   string         `json:"created_at" db:"created_at"`
	CompletedAt string         `json:"completed_at" db:"completed_at"`
	Items       []*ItemInOrder `json:"items"`
}

type ItemInOrder struct {
	ItemId  string `json:"item_id" db:"item_id"`
	OrderId string `db:"order_id"`
	Count   int    `json:"count" db:"count"`
	Price   int    `json:"price" db:"price"`
}
