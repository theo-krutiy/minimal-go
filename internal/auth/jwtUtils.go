package auth

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

func issueNewJWT(secret []byte) (string, error) {
	claims := jwt.MapClaims{
		"exp": jwt.NewNumericDate(time.Now().Add(24 * time.Hour)),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	ss, err := token.SignedString(secret)
	if err != nil {
		return "", err
	}
	return ss, nil
}

func validateJWT(tokenString string, secret []byte) error {
	var kf = func(token *jwt.Token) (interface{}, error) { return secret, nil }
	_, err := jwt.Parse(tokenString, kf, jwt.WithExpirationRequired(), jwt.WithValidMethods([]string{"HS256"}))
	switch {
	case errors.Is(err, jwt.ErrTokenExpired):
		return errors.New("token expired")
	case err != nil:
		return errors.New("malformed token")
	default:
		return nil
	}
}
