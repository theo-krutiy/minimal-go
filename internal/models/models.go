package models

type UserInDatabase struct {
	Id           string
	Login        string
	PasswordHash []byte
}

type ItemInDatabase struct {
	Id             string
	Name           string
	CountAvailable int
	PriceInteger   int
	PriceDecimal   int
}
