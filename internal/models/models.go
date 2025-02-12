package models

type UserInDatabase struct {
	Id           string
	Login        string
	PasswordHash []byte
}

type ItemInDatabase struct {
	Id             string `json:"id"`
	Name           string `json:"name"`
	CountAvailable int    `json:"count_available"`
	PriceInteger   int    `json:"price_integer"`
	PriceDecimal   int    `json:"price_decimal"`
}
