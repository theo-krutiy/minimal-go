package auth

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

func newClaims() jwt.Claims {
	return jwt.RegisteredClaims{
		ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)),
	}
}

func issueNewJWT(secret []byte) (string, error) {
	claims := newClaims()
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	ss, err := token.SignedString(secret)
	if err != nil {
		return "", err
	}
	return ss, nil
}

func validateJWT(tokenString string, secret []byte) error {
	var kf = func(token *jwt.Token) (interface{}, error) { return secret, nil }
	token, err := jwt.Parse(tokenString, kf, jwt.WithExpirationRequired())
	switch {
	case token.Valid:
		return nil
	case errors.Is(err, jwt.ErrTokenExpired):
		return errors.New("token expired")
	default:
		return errors.New("malformed token")
	}
}
