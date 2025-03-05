package token

import (
	"github.com/golang-jwt/jwt/v4"
)

type CreateTokenPayload struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type CreateTokenResponse struct {
	Token string `json:"token"`
}

type Generator interface {
	GenerateJwtToken(secret string, claims jwt.Claims) (string, error)
}

type CustomToken struct {
}

func (c *CustomToken) GenerateJwtToken(secret string, claims jwt.Claims) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString([]byte(secret))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}
