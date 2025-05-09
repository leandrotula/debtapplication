package token

import (
	"github.com/golang-jwt/jwt/v4"
)

type TokenErrorResponse struct {
	Error string `json:"error"`
}

type CreateTokenPayload struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type CreateTokenResponse struct {
	Token string `json:"token"`
}

func NewTokenResponse(token string) *CreateTokenResponse {
	return &CreateTokenResponse{
		token,
	}
}

func NewCustomTokenError(error string) *TokenErrorResponse {
	return &TokenErrorResponse{
		error,
	}
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
